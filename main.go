package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"./session"
)

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
	token    interface{}
}

// ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
		/** @ToDo token **/
	})
	token, _ := session.GetFlashSession(w, r)
	data := map[string]interface{}{
		"Host":  r.Host,
		"Token": token,
	}
	//	if authCookie, err := r.Cookie("auth"); err == nil {
	//		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	//	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈します

	/**
	 * @ToDo 作成するページ
	 * ・初期化ページ　ログインページにリダイレクトをした時、データベースが生成されていなければinitを表示する
	 * ・ログインページ　ログインしていなければ必ずここ
	 * ・法人名
	 * ・圃場名
	 * ・土壌分析登録ページ
	 * ・APIのページ
	 */

	//ログインしているか調べる
	http.Handle("/admin", MustAuth(&templateHandler{filename: "admin.html"}))

	//ログインページを開くときは常にAdministratorのテーブルがあるか調べる
	http.Handle("/login", CheckDB(&templateHandler{filename: "login.html"}))
	http.Handle("/init", &initHandler{filename: "init.html"})

	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
