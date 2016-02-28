package session

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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

func Get(r *http.Request, session_key string) (s *sessions.Session) {
	s, _ = store.Get(r, session_key)
	return s
}

func Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) {
	s.Save(r, w)
}

func GetFlashSession(w http.ResponseWriter, r *http.Request) (token string, go_token string) {
	var n uint64

	binary.Read(rand.Reader, binary.LittleEndian, &n)
	token = strconv.FormatUint(n, 36)

	//フラッシュセッションの値を取り出す
	s, _ := store.Get(r, "soilapp-token")
	if flashes := s.Flashes(); len(flashes) > 0 {
		go_token, _ = flashes[0].(string)
	}

	s.AddFlash(token)
	s.Save(r, w)

	return token, go_token
}

type SessionConfig struct {
	Key string
}

func getSessionConfig() SessionConfig {
	var sconf SessionConfig

	p, _ := os.Getwd()
	if strings.Index(p, "\\") > 0 {
		p = strings.Replace(p, "\\", "/", -1)
	}
	p = strings.Replace(p, "/session", "", 1)

	jsonString, err := ioutil.ReadFile(p + "/conf/session.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonString, &sconf)
	if err != nil {
		panic(err)
	}

	return sconf
}
