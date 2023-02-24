package middleware

import (
	"context"
	m "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/gtoxlili/give-advice/domain/response"
	log "github.com/sirupsen/logrus"
	"net/http"
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

// CheckOpenAIToken 检验 openAI Token 是否有效
// TODO 以后再说
func CheckOpenAIToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// 将 Token 放入上下文
		r.WithContext(context.WithValue(r.Context(), "token", r.Header.Get("OpenAI-Auth-Key")))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
