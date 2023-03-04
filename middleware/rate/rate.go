package rate

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/gtoxlili/advice-hub/domain/response"
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

func ExceededHandler(w http.ResponseWriter, r *http.Request) {
	retryStr := w.Header().Get("Retry-After")
	retrySeconds, _ := strconv.Atoi(retryStr)
	render.JSON(w, r, response.Fail(429, fmt.Sprintf("Rate limit exceeded, retry in %d minutes", retrySeconds/60)))
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
