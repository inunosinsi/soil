package view

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type jsonHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func NewJsonHandler(filename string) jsonHandler {
	return jsonHandler{filename: filename}
}

func (h *jsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.once.Do(func() {
		h.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				h.filename)))
	})

	h.templ.Execute(w, nil)
}
