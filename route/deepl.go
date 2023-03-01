package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gtoxlili/give-advice/common/validate"
	"github.com/gtoxlili/give-advice/components/deepl"
	"github.com/gtoxlili/give-advice/domain/response"
)

type translateBody struct {
	Content string `json:"content" validate:"required"`
}

func (t *translateBody) Bind(r *http.Request) error {
	return validate.Struct(t)
}

func Deepl(r chi.Router) {
	r.Post("/translate/{target}", translate)
}

func translate(w http.ResponseWriter, r *http.Request) {
	target := chi.URLParam(r, "target")
	if !deepl.IsValidLang(target) {
		render.JSON(w, r, response.Fail(400, "invalid target"))
		return
	}

	options := make([]deepl.Option, 1)
	if r.URL.Query().Get("isHtml") != "" {
		options = append(options, deepl.WithHTML())
	}

	body := &translateBody{}
	if err := render.Bind(r, body); err != nil {
		render.JSON(w, r, response.Fail(400, err.Error()))
		return
	}
	text, err := deepl.Translate(r.Context(), body.Content, target, options...)
	if err != nil {
		render.JSON(w, r, response.Fail(500, err.Error()))
		return
	}
	render.JSON(w, r, response.Ok(text))
}
