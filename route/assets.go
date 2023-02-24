package route

import (
	"embed"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
)

func Assets(fs embed.FS) func(r chi.Router) {
	return func(r chi.Router) {
		h := assets(fs)
		r.Handle("/", h)
		r.Handle("/index.html", h)
		r.Handle("/icon.ico", h)
		r.Handle("/sw.js", h)
		r.Handle("/manifest.webmanifest", h)
		r.Handle("/workbox-*", h)
		r.Handle("/assets/*", h)
	}
}

func assets(assets embed.FS) http.Handler {
	sub, err := fs.Sub(assets, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}
	return http.FileServer(http.FS(sub))
}
