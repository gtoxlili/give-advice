package rate

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/gtoxlili/give-advice/domain/response"
)

func Token(r *http.Request) (string, error) {
	// TODO 暂且通过 IP 作为用户令牌
	return KeyByRealIP(r)
}

func LimitKeyFunc(r *http.Request) (string, error) {
	if token, ok := r.Context().Value("Token").(string); ok {
		return token, nil
	}
	return "", fmt.Errorf("token not found")
}

func ExceededHandler(retry time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 整数
		render.JSON(w, r, response.Fail(429, fmt.Sprintf("Rate limit exceeded, retry in %.0f seconds", retry.Seconds())))
	}
}

func TokenIntoCtx(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := Token(r)
		if err != nil {
			render.JSON(w, r, response.Fail(400, fmt.Sprintf("invalid token: %s", err)))
			return
		}
		ctx := context.WithValue(r.Context(), "Token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
