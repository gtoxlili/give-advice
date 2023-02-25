package route

import (
	"net/http"
	"sync/atomic"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gtoxlili/give-advice/domain/response"
)

func Info(r chi.Router) {
	// 使用人次
	r.Get("/useCount", useCount)
}

// Count 使用人次
// TODO 先做简单处理 拿个原子存一下
var Count atomic.Int64

func useCount(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, response.Ok(
		render.M{"count": Count.Load()},
	))
}
