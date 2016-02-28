package view

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"../session"
)

type orgHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func NewOrgHandler(filename string) orgHandler {
	return orgHandler{filename: filename}
}

func (h *orgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {
			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {

			}
		}
	}

	h.once.Do(func() {
		h.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				h.filename)))
	})
	token, _ := session.GetFlashSession(w, r)
	data := map[string]interface{}{
		"Token": token,
	}

	h.templ.Execute(w, data)
}
