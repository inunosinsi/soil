package session

import (
	"io/ioutil"
	"net/http"

	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"strconv"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	sconf := getSessionConfig()
	store = sessions.NewCookieStore([]byte(sconf.Key))
}

func Get(r *http.Request, session_key string) (session *sessions.Session) {
	session, err := store.Get(r, session_key)
	if err != nil {
		panic(err)
	}

	return session
}

func Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) {
	session.Save(r, w)
}

func GetFlashSession(w http.ResponseWriter, r *http.Request) (token string, go_token string) {
	var n uint64

	binary.Read(rand.Reader, binary.LittleEndian, &n)
	token = strconv.FormatUint(n, 36)

	//フラッシュセッションの値を取り出す
	session, _ := store.Get(r, "soilapp-token")
	if flashes := session.Flashes(); len(flashes) > 0 {
		go_token, _ = flashes[0].(string)
	}

	session.AddFlash(token)
	session.Save(r, w)

	return token, go_token
}

type SessionConfig struct {
	Key string
}

func getSessionConfig() SessionConfig {
	var sconf SessionConfig

	jsonString, err := ioutil.ReadFile("../conf/session.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonString, &sconf)
	if err != nil {
		panic(err)
	}

	return sconf
}
