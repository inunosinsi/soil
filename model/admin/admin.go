package admin

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mholt/binding"

	"../../goy/goydb"
)

type Admin struct {
	Id       int
	LoginId  string
	Password string
}

func NewAdmin() Admin {
	return Admin{}
}

func (a *Admin) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&a.Id:       "id",
		&a.LoginId:  "login_id",
		&a.Password: "password",
	}
}

func (a *Admin) TableName() string {
	return "Administrator"
}

func Check() bool {
	db := goydb.Conn()
	defer db.Close()

	res, err := db.Query("SELECT id FROM Administrator LIMIT 1")
	if err != nil {
		db.Close()
		return false
	}

	for res.Next() {
		var id int
		err = res.Scan(&id)
		if err != nil {
			db.Close()
			return false
		}

		//データがあればtrueを返す
		if id > 0 {
			break
		}
	}

	return true
}

func Insert(a *Admin) int64 {
	var dbs goydb.Goydb = a
	id, err := goydb.Insert(dbs)
	if err != nil {
		panic(err)
	}

	return id
}

func GetPasswordHashByLoginId(loginId string) string {
	var passHash string

	db := goydb.Conn()
	defer db.Close()

	stmt, err := db.Prepare("SELECT password FROM Administrator WHERE login_id = ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(loginId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&passHash)
		if err != nil {
			panic(err)
		}
	}

	return passHash
}
