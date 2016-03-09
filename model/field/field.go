package field

import (
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mholt/binding"

	"../../goy/goydb"
)

type Field struct {
	Id    int
	Name  string
	OrgId int
}

func NewField() Field {
	return Field{}
}

func (f *Field) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Id:    "id",
		&f.Name:  "name",
		&f.OrgId: "org_id",
	}
}

func (f *Field) TableName() string {
	return "Field"
}

func Insert(field *Field) int64 {
	var dbs goydb.Goydb = field
	id, err := goydb.Insert(dbs)
	if err != nil {
		panic(err)
	}

	return id
}

func Get(limit int) *[]Field {
	db := goydb.Conn()
	defer db.Close()

	lim := strconv.Itoa(limit)

	rows, err := db.Query("SELECT * FROM Field LIMIT " + lim)
	if err != nil {
		panic(err.Error())
	}

	list := make([]Field, 0)

	for rows.Next() {
		var id int
		var name string
		var orgId int
		err = rows.Scan(&id, &name, &orgId)
		if err != nil {
			panic(err)
		}
		if id > 0 {
			list = append(list, Field{id, name, orgId})
		}
	}

	return &list
}

func GetById(fieldId int) *Field {
	var f Field
	
	db := goydb.Conn()
	defer db.Close()
	
	stmt, err := db.Prepare("SELECT * FROM Field WHERE id = ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(fieldId)
	if err != nil {
		panic(err)
	}
		
	for rows.Next() {
		var id int
		var name string
		var orgId int
		err = rows.Scan(&id, &name, &orgId)
		if err != nil {
			panic(err)
		}
		f = Field{id, name, orgId}
	}
	
	return &f
}

func GetByOrgId(orgId int) *[]Field {
	db := goydb.Conn()
	defer db.Close()
	
	stmt, err := db.Prepare("SELECT * FROM Field WHERE org_id = ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(orgId)
	if err != nil {
		panic(err)
	}
	
	list := make([]Field, 0)
	
	for rows.Next() {
		var id int
		var name string
		var orgId int
		err = rows.Scan(&id, &name, &orgId)
		if err != nil {
			panic(err)
		}
		if id > 0 {
			list = append(list, Field{id, name, orgId})
		}
	}

	return &list
}
