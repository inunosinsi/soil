package view

import (
	"html"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"../model/field"
	"../model/org"
	"../session"
)

type fieldHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func NewFieldHandler(filename string) fieldHandler {
	return fieldHandler{filename: filename}
}

func (h *fieldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := strings.Replace(r.URL.Path, "/org/", "", 1)
	orgId, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	if r.Method == "POST" {
		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {
			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				if name := html.EscapeString(r.FormValue("name")); len(name) > 0 {
					id := field.Insert(name, orgId)
					oid := strconv.Itoa(orgId)
					if id > 0 {
						w.Header().Set("Location", "/org/"+oid+"?successed")
						w.WriteHeader(http.StatusTemporaryRedirect)
					} else {
						w.Header().Set("Location", "/org/"+oid+"?error")
						w.WriteHeader(http.StatusTemporaryRedirect)
					}
				}
			}
		}

	}

	h.once.Do(func() {
		h.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				h.filename)))
	})

	org := org.GetById(orgId)
	fields := field.Get(10)

	token, _ := session.GetFlashSession(w, r)
	data := map[string]interface{}{
		"Token":  token,
		"Org":    org,
		"Fields": fields,
	}

	h.templ.Execute(w, data)
}
