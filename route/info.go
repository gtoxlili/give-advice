package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gtoxlili/advice-hub/domain/response"
	"github.com/gtoxlili/advice-hub/repository/redis"
)

func Info(r chi.Router) {
	// 使用人次
	r.Get("/useCount", useCount)
}

func useCount(w http.ResponseWriter, r *http.Request) {
	count, err := redis.Default().GetVisitCount(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, response.Ok(
		render.M{"count": count},
	))
}
