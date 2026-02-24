package server

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func prodSetup(r *chi.Mux) {
	// Production serve
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		// Block access to .vite directory
		if strings.Contains(r.URL.Path, ".vite") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var data []byte
		target, err := getTargetFile("./internal/server/frontend/dist", r.URL.Path)
		if err == nil {
			data, err = io.ReadAll(target)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", getFileContentType(filepath.Ext(r.URL.Path), data))
			_, _ = w.Write(data)
			return
		} else if filepath.Ext(r.URL.Path) != "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		executeProdIndex(w)
	})
}

func devSetup(r *chi.Mux) {
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		executeDevIndex(w)
	})
}

func setupFrontend(r *chi.Mux, dev bool) {
	if !dev {
		prodSetup(r)
		return
	}
	devSetup(r)
}
