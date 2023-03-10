package main

import (
	"embed"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/gtoxlili/advice-hub/log"
	"github.com/gtoxlili/advice-hub/middleware"
	"github.com/gtoxlili/advice-hub/middleware/rate"
	"github.com/gtoxlili/advice-hub/route"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//go:embed frontend/dist/*
var assets embed.FS

func main() {
	r := chi.NewRouter()

	// 通用中间件
	r.Use(middleware.Recover)
	r.Use(middleware.Cors("*"))

	// 静态资源
	r.Group(route.Assets(assets))

	r.Route("/api", func(r chi.Router) {
		r.Route("/info", route.Info)
		r.Group(func(r chi.Router) {
			// r.Use(auth.Authenticated)
			r.Use(rate.TokenIntoCtx)
			r.Use(middleware.Logger)
			r.Group(route.OpenAI)
			r.Route("/deepl", route.Deepl)
		})
	})

	h2s := h2c.NewHandler(r, &http2.Server{})
	l, err := net.Listen("tcp", "127.0.0.1:7458")
	if err != nil {
		log.Fatalln(err)
	}
	log.Infof("Server listening at %s", l.Addr())
	if err = http.Serve(l, h2s); err != nil {
		log.Fatalln(err)
	}

}
