package main

import (
	"html"
	"html/template"
	"net/http"
	"path/filepath"

	"./login"
	"./model/admin"
	"./session"
)

type initHandler struct {
	filename string
	t        templateHandler
}

func (h *initHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {

			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				password := html.EscapeString(r.FormValue("password"))
				confirm := html.EscapeString(r.FormValue("confirm"))

				//入力した内容が一致した時
				if password == confirm {
					loginId := html.EscapeString(r.FormValue("login_id"))
					hash := login.CreateHashString(password)
					id := admin.Insert(loginId, hash)
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

	h.t.templ = template.Must(template.ParseFiles(filepath.Join("templates", h.filename)))
	token, _ := session.GetFlashSession(w, r)
	data := map[string]interface{}{
		"Token": token,
	}

	h.t.templ.Execute(w, data)
}
