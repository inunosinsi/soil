package view

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"../model/analysis"
	"../model/field"
	"../session"

	"github.com/mholt/binding"
)

type analysisHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func NewAnalysisHandler(filename string) analysisHandler {
	return analysisHandler{filename: filename}
}

func (h *analysisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := strings.Replace(r.URL.Path, "/analysis/", "", 1)
	aId, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	if r.Method == "POST" && len(r.URL.RawQuery) == 0 {
		aid := strconv.Itoa(aId)

		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {
			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				a := analysis.NewAnalysis()
				err := binding.Bind(r, &a)
				if err != nil {
					panic(err)
				}

				analysis.Update(&a)
				w.Header().Set("Location", "/analysis/"+aid+"?successed")
				w.WriteHeader(http.StatusTemporaryRedirect)
			}
		}

		w.Header().Set("Location", "/analysis/"+aid+"?error")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	a := analysis.GetById(aId)

	h.once.Do(func() {
		h.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				h.filename)))
	})

	f := field.GetById(a.FieldId)
	token, _ := session.GetFlashSession(w, r)
	data := map[string]interface{}{
		"Token":    token,
		"Analysis": a,
		"Field":    f,
	}

	h.templ.Execute(w, data)
}
