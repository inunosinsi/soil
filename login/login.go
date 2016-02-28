package login

import (
	"encoding/base32"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/sha3"

	"../session"
	"../model/admin"
)

func CheckPassword(loginId string, password string) bool {
	passwordHash := admin.GetPasswordHashByLoginId(loginId)
	salt, oldHash := GetSaltAndHashByPasswordHash(passwordHash)

	if newHash := CulcHash(password, salt); newHash == oldHash {
		return true
	} else {
		return false
	}
}

func Login(w http.ResponseWriter, r *http.Request, loginId string) {
	s := session.Get(r, "soilapp-login")
	s.Values["loginId"] = loginId
	session.Save(r, w, s)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	s := session.Get(r, "soilapp-login")
	s.Values["loginId"] = nil
	session.Save(r, w, s)
}

func IsLoggedIn(r *http.Request) bool {
	s := session.Get(r, "soilapp-login")
	if loginId, _ := s.Values["loginId"].(string); len(loginId) > 0 {
		return true
	} else {
		return false
	}
}

//パスワードカラムに突っ込む規格/ソルト/ハッシュ値の値を生成する
func CreateHashString(password string) string {
	rand.Seed(time.Now().UnixNano())
	salt := strconv.Itoa(rand.Intn(999))

	hash := CulcHash(password, salt)

	return "sha3/" + salt + "/" + hash
}

func GetSaltAndHashByPasswordHash(passwordHash string) (salt string, hash string) {
	//配列 0:暗号化の規格、1:ソルト、2:ハッシュ値
	values := strings.Split(passwordHash, "/")
	return values[1], values[2]
}

//パスワードとソルトからハッシュ値を計算する
func CulcHash(password string, salt string) string {
	sh3 := sha3.New256()
	io.WriteString(sh3, password+salt)

	// Conversion to base32
	return strings.ToLower(base32.HexEncoding.EncodeToString(sh3.Sum(nil)))
}
