package view

import (
	"html"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"../login"
	"../model/analysis"
	"../model/field"
	"../model/org"
	"../session"

	"github.com/mholt/binding"
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

	if r.Method == "POST" && len(r.URL.RawQuery) == 0 {
		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {
			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				if name := html.EscapeString(r.FormValue("name")); len(name) > 0 {
					f := field.NewField()
					err := binding.Bind(r, &f)
					if err != nil {
						panic(err)
					}

					f.OrgId = orgId
					id := field.Insert(&f)
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
	fields := field.GetByOrgId(orgId)

	token, _ := session.GetFlashSession(w, r)
	data := map[string]interface{}{
		"Token":  token,
		"Org":    org,
		"Fields": fields,
	}

	h.templ.Execute(w, data)
}

/** ここから詳細 **/

type fieldDetailHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func NewFieldDetailHandler(filename string) fieldDetailHandler {
	return fieldDetailHandler{filename: filename}
}

func (h *fieldDetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := strings.Replace(r.URL.Path, "/field/", "", 1)
	fieldId, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	field := field.GetById(fieldId)

	if r.Method == "POST" && len(r.URL.RawQuery) == 0 {

		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {
			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				if date := html.EscapeString(r.FormValue("analysis_date")); len(date) > 0 {
					a := analysis.NewAnalysis()
					err := binding.Bind(r, &a)
					if err != nil {
						panic(err)
					}

					a.FieldId = fieldId

					//フィールドKEYの登録 フィールド名とフィールドIDでハッシュを作る
					fid := strconv.Itoa(fieldId)
					a.FieldKey = login.CulcHash(field.Name, fid)

					id := analysis.Insert(&a)
					if id > 0 {
						w.Header().Set("Location", "/field/"+fid+"?successed")
						w.WriteHeader(http.StatusTemporaryRedirect)
					} else {
						w.Header().Set("Location", "/field/"+fid+"?error")
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

	analysis := analysis.GetByFieldId(fieldId)

	token, _ := session.GetFlashSession(w, r)
	data := map[string]interface{}{
		"Token":    token,
		"Field":    field,
		"Analysis": analysis,
	}

	h.templ.Execute(w, data)
}
