package middleware

import (
	"net/http"

	m "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/gtoxlili/give-advice/domain/response"
	log "github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return m.RequestLogger(&m.DefaultLogFormatter{
		Logger: log.StandardLogger(),
	})(next)
}

func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				render.JSON(w, r, response.FailWith(err.(error)))
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Cors(origin ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return cors.Handler(cors.Options{
			AllowedOrigins:   origin,
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowedHeaders:   []string{"Content-Type", "Authorization", "OpenAI-Auth-Key"},
			MaxAge:           300,
			AllowCredentials: true,
		})(next)
	}
}

// StripBearer 去掉 OpenAI-Auth-Key 的 Bearer 前缀
func StripBearer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("OpenAI-Auth-Key"); auth != "" {
			r.Header.Set("OpenAI-Auth-Key", auth[7:])
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
