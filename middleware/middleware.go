package middleware

import (
	"net/http"

	m "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/gtoxlili/advice-hub/domain/response"
	"github.com/gtoxlili/advice-hub/repository/redis"
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

func CustomOpenAIToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if auth := r.Header.Get("OpenAI-Auth-Key"); auth != "" {
			token, _ := r.Context().Value("Token").(string)
			log.Infof("user [%s] use custom token", token)
			r.Header.Set("OpenAI-Auth-Key", auth[7:])
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// IncrVisitCount 记录访问次数
func IncrVisitCount(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if token, ok := r.Context().Value("Token").(string); ok {
			go func() {
				val, _ := redis.Default().IncrVisitCount(r.Context(), token)
				log.Infof("[%s] visit count: %d", token, val)
			}()
			next.ServeHTTP(w, r)
		} else {
			render.JSON(w, r, response.Fail(401, "token not found"))
		}
	}
	return http.HandlerFunc(fn)
}
