package route

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
	"github.com/gtoxlili/give-advice/common/cache"
	"github.com/gtoxlili/give-advice/common/unsafe"
	"github.com/gtoxlili/give-advice/common/validate"
	"github.com/gtoxlili/give-advice/domain/response"
	m "github.com/gtoxlili/give-advice/middleware"
	"github.com/gtoxlili/give-advice/middleware/rate"
	"github.com/gtoxlili/give-advice/openai"
	"github.com/jaevor/go-nanoid"
)

func OpenAI(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(middleware.RouteHeaders().
			Route("OpenAI-Auth-Key", "bearer *", m.StripBearer).
			RouteDefault(
				httprate.Limit(2, time.Minute,
					httprate.WithKeyFuncs(rate.LimitKeyFunc),
					httprate.WithLimitHandler(rate.ExceededHandler(time.Minute)),
				),
			).Handler,
		)
		r.Use(m.IncrVisitCount)
		r.Post("/register/{type}", register)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(middleware.ThrottleBacklog(32, 2048, time.Minute))
		r.Get("/inquiry/{id}", inquiry)
	})
}

type registerBody struct {
	OpenAIToken string
	Type        string
	Noun        string `json:"noun" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (b *registerBody) Bind(_ *http.Request) error {
	return validate.Struct(b)
}

var nanoIdGenerator, _ = nanoid.Standard(21)
var registerCache = cache.New(cache.WithAge(120), cache.WithSize(512))

func register(w http.ResponseWriter, r *http.Request) {
	typ := chi.URLParam(r, "type")
	if openai.Factory(typ) == nil {
		render.JSON(w, r, response.Fail(400, "invalid type"))
		return
	}

	body := &registerBody{
		OpenAIToken: r.Header.Get("OpenAI-Auth-Key"),
		Type:        typ,
	}
	if err := render.Bind(r, body); err != nil {
		render.JSON(w, r, response.Fail(400, err.Error()))
		return
	}

	nanoID := nanoIdGenerator()
	registerCache.Set(nanoID, body)
	render.JSON(w, r, response.Ok(render.M{"id": nanoID}))
}

func inquiry(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")
	// nginx sse 支持
	w.Header().Set("X-Accel-Buffering", "no")

	id := chi.URLParam(r, "id")
	obj, ok := registerCache.Get(id)
	defer registerCache.Delete(id)
	if !ok {
		msg := fmt.Sprintf("event: error\ndata: %s\n\n", "invalid id")
		_, err := w.Write(unsafe.ToByte(msg))
		if err == nil {
			w.(http.Flusher).Flush()
		}
		return
	}
	body := obj.(*registerBody)
	// 将 可能存在的 Token 放入上下文
	r = r.WithContext(context.WithValue(r.Context(), "OpenAI-Auth-Key", body.OpenAIToken))
	ch := openai.Factory(body.Type)(r.Context(), body.Noun, body.Description)
	modCh := openai.ModerationChan(r.Context(), body.Noun+"\n"+body.Description)

	for msg, flag := "", true; flag; {
		flag = false
		select {
		case mod := <-modCh:
			msg = fmt.Sprintf("event: error\ndata: %s\n\n", mod)
		case result, more := <-ch:
			if !more {
				msg = "event: end\ndata: \n\n"
			} else {
				if result.Err != nil {
					msg = fmt.Sprintf("event: error\ndata: %s\n\n", result.Err.Error())
				} else {
					msg = fmt.Sprintf("data: %s\n\n", result.Val)
					flag = true
				}
			}
		}
		_, err := w.Write(unsafe.ToByte(msg))
		if err != nil {
			return
		}
		w.(http.Flusher).Flush()
	}
}
