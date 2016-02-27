package login

import "testing"

func TestCreateHashString(t *testing.T) {
	pw := "password"
	hash := CreateHashString(pw)

	if hash == "" {
		t.Error("ハッシュ化されず空文字が返ってきました")
	}

	if pw == hash {
		t.Error("指定したパスワードがハッシュ化されていません")
	}

	con := "confirm"
	newHash := CreateHashString(con)

	if hash == newHash {
		t.Error("違う文字列で同じハッシュ値が生成されました")
	}
}

func TestCulcHash(t *testing.T) {
	pw := "password"
	hash := CulcHash(pw, "123")

	if hash == "" {
		t.Error("ハッシュ化されず空文字が返ってきました")
	}

	if hash == pw {
		t.Error("指定されたパスワードがハッシュ化されていません")
	}

	newHash := CulcHash(pw, "456")

	if hash == newHash {
		t.Error("ハッシュ化の際にソルトが効いていません")
	}
}
