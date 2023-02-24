package main

import (
	"embed"
	"github.com/go-chi/chi/v5"
	_ "github.com/gtoxlili/give-advice/log"
	"github.com/gtoxlili/give-advice/middleware"
	"github.com/gtoxlili/give-advice/route"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

//go:embed frontend/dist/*
var assets embed.FS

func main() {
	r := chi.NewRouter()

	// 通用中间件
	r.Use(middleware.Logger)
	r.Use(middleware.Recover)
	r.Use(middleware.Cors("*"))

	// 静态资源
	r.Group(route.Assets(assets))

	r.Route("/api", func(r chi.Router) {
		// r.Use(auth.Authenticated)
		r.Group(route.OpenAI)
		r.Route("/deepl", route.Deepl)
		r.Route("/info", route.Info)
	})

	l, err := net.Listen("tcp", "127.0.0.1:7458")
	if err != nil {
		log.Fatalln(err)
	}
	log.Infof("Server listening at %s", l.Addr())
	if err = http.Serve(l, r); err != nil {
		log.Fatalln(err)
	}

}
