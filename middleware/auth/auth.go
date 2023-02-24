package auth

import (
	"net/http"
)

type Rule uint8

const (
	Anonymous Rule = iota
	NonMember
	Member
	Admin
)

func Authentication(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// todo 验证token
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func HasRule(rule Rule) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO
			next.ServeHTTP(w, r)
		})
	}
}
