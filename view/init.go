package view

import (
	"html"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"../login"
	"../model/admin"
	"../session"
	
	"github.com/mholt/binding"
)

func NewInitHandler(filename string) initHandler {
	return initHandler{filename: filename}
}

type initHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (h *initHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && len(r.URL.RawQuery) == 0 {
		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {

			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				password := html.EscapeString(r.FormValue("password"))
				confirm := html.EscapeString(r.FormValue("confirm"))

				//入力した内容が一致した時
				if password == confirm {
					a := admin.NewAdmin()
					err := binding.Bind(r, &a)
					if err != nil {
						panic(err)
					}
					
					a.Password = login.CreateHashString(password)
					id := admin.Insert(&a)
					
					if id > 0 {
						//ログインページへ飛ぶ
						w.Header().Set("Location", "/login")
						w.WriteHeader(http.StatusTemporaryRedirect)
					} else {
						w.Header().Set("Location", "/init?error")
						w.WriteHeader(http.StatusTemporaryRedirect)
					}
				}
			}
		}
	}

	//既に管理者がいないかチェック
	if isAdmin := admin.Check(); isAdmin {
		//ログインページへ飛ぶ
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
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
