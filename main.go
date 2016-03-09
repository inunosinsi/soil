package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"./session"
	"./view"
)

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
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
	//外部cssや外部js用
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	//管理画面トップ
	orgHandler := view.NewOrgHandler("admin.html")
	http.Handle("/admin", MustAuth(&orgHandler))

	fieldHandler := view.NewFieldHandler("org.html")
	http.Handle("/org/", MustAuth(&fieldHandler)) //URLがつづく場合は末尾にスラッシュ

	fieldDetailHandler := view.NewFieldDetailHandler("field.html")
	http.Handle("/field/", MustAuth(&fieldDetailHandler)) //URLがつづく場合は末尾にスラッシュ
	
	analysisHandler := view.NewAnalysisHandler("analysis.html")
	http.Handle("/analysis/", MustAuth(&analysisHandler)) //URLがつづく場合は末尾にスラッシュ
	
	graphHandler := view.NewGraphHandler("graph.html")
	http.Handle("/graph/", &graphHandler) //URLがつづく場合は末尾にスラッシュ

	//ログインページを開くときは常にAdministratorのテーブルがあるか調べる
	http.Handle("/login", view.CheckDB(&templateHandler{filename: "login.html"}))

	http.Handle("/logout", view.Logout(&templateHandler{filename: "logout.html"}))

	initHandler := view.NewInitHandler("init.html")
	http.Handle("/init", &initHandler)

	jsonHandler := view.NewJsonHandler()
	http.Handle("/call.json", &jsonHandler)

	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
