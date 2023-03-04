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
	"github.com/gtoxlili/advice-hub/common/cache"
	"github.com/gtoxlili/advice-hub/common/unsafe"
	"github.com/gtoxlili/advice-hub/components/openai"
	"github.com/gtoxlili/advice-hub/domain/response"
	m "github.com/gtoxlili/advice-hub/middleware"
	"github.com/gtoxlili/advice-hub/middleware/rate"
	"github.com/jaevor/go-nanoid"
)

func OpenAI(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(middleware.RouteHeaders().
			Route("OpenAI-Auth-Key", "bearer *", m.CustomOpenAIToken).
			RouteDefault(
				httprate.Limit(10, 10*time.Minute,
					httprate.WithKeyFuncs(rate.LimitKeyFunc),
					httprate.WithLimitHandler(rate.ExceededHandler),
				),
			).Handler,
		)
		r.Use(m.IncrVisitCount)
		r.Post("/register/{type}", register)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(middleware.ThrottleBacklog(32, 2048, time.Minute))
		r.Get("/inquiry/{id}", completion)
	})
}

type registerDto struct {
	OpenAIToken string
	Data        openai.Adaptor
}

var nanoIdGenerator, _ = nanoid.Standard(21)
var registerCache = cache.New(cache.WithAge(120), cache.WithSize(512))

func register(w http.ResponseWriter, r *http.Request) {
	typ := chi.URLParam(r, "type")
	adaptor := openai.Factory(typ)
	if adaptor == nil {
		render.JSON(w, r, response.Fail(400, "invalid type"))
		return
	}
	if err := render.Bind(r, adaptor); err != nil {
		render.JSON(w, r, response.Fail(400, err.Error()))
		return
	}
	nanoID := nanoIdGenerator()
	registerCache.Set(nanoID, &registerDto{
		OpenAIToken: r.Header.Get("OpenAI-Auth-Key"),
		Data:        adaptor,
	})
	render.JSON(w, r, response.Ok(render.M{"id": nanoID}))
}

func completion(w http.ResponseWriter, r *http.Request) {

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
	body := obj.(*registerDto)
	// 将 可能存在的 Token 放入上下文
	r = r.WithContext(context.WithValue(r.Context(), "OpenAI-Auth-Key", body.OpenAIToken))
	modCh := openai.ModerationChan(r.Context(), body.Data.Review())
	ch := body.Data.Completion(r.Context())

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
