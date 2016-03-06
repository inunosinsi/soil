package view

import (
	"html"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"../model/org"
	"../session"

	"github.com/mholt/binding"
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
	if r.Method == "POST" && len(r.URL.RawQuery) == 0 {
		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {
			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				if name := html.EscapeString(r.FormValue("name")); len(name) > 0 {
					o := org.NewOrg()
					err := binding.Bind(r, &o)
					if err != nil {
						panic(err)
					}

					id := org.Insert(&o)
					if id > 0 {
						w.Header().Set("Location", "/admin?successed")
						w.WriteHeader(http.StatusTemporaryRedirect)
					}
				}
			}
		}

		w.Header().Set("Location", "/admin?error")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	h.once.Do(func() {
		h.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				h.filename)))
	})

	token, _ := session.GetFlashSession(w, r)
	orgs := org.Get(10)
	data := map[string]interface{}{
		"Token": token,
		"Orgs":  orgs,
	}

	h.templ.Execute(w, data)
}
