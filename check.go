package main

import (
	"html"
	"net/http"

	"./login"
	"./model/admin"
	"./session"
)

type checkHandler struct {
	next http.Handler
}

func (h *checkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if post_token := r.FormValue("go_token"); len(post_token) > 0 {

			if _, go_token := session.GetFlashSession(w, r); post_token == go_token {
				loginId := html.EscapeString(r.FormValue("login_id"))
				password := html.EscapeString(r.FormValue("password"))

				//ログイン
				if login.CheckPassword(loginId, password) {
					login.Login(w, r, loginId)
					http.Redirect(w, r, "/admin", http.StatusFound)
				} else {
					http.Redirect(w, r, "/login?error", http.StatusFound)
				}
			}
		}
	}

	isExisted := admin.Check()

	//初期化フラグがtrueの場合は初期化ページへ
	if !isExisted {
		//初期化ページへ飛ぶ
		w.Header().Set("Location", "/init")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	if isLoggedIn := login.IsLoggedIn(r); isLoggedIn {
		// 未認証
		w.Header().Set("Location", "/admin")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	// 成功。ラップされたハンドラを呼び出します
	h.next.ServeHTTP(w, r)
}
func CheckDB(handler http.Handler) http.Handler {
	return &checkHandler{next: handler}
}
