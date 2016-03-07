package analysis

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mholt/binding"

	"../../goy/goydb"
)

type Analysis struct {
	Id       int
	FieldId  int
	FieldKey string
	Ph       float64
	Phk      float64
	Ec       float64
	Php      float64
	Eofph    float64
	K        float64
	Dk       float64
	Ca       float64
	Dca      float64
	Mg       float64
	Dmg      float64
	Cec      int
	Dcec     float64
	Capermg  float64
	Mgperk   float64
}

func NewAnalysis() Analysis {
	return Analysis{}
}

func (a *Analysis) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&a.Id:       "id",
		&a.FieldId:  "field_id",
		&a.FieldKey: "field_key",
		&a.Ph:       "ph",
		&a.Phk:      "phk",
		&a.Ec:       "ec",
		&a.Php:      "php",
		&a.Eofph:    "eofph",
		&a.K:        "k",
		&a.Dk:       "dk",
		&a.Ca:       "ca",
		&a.Dca:      "dca",
		&a.Mg:       "mg",
		&a.Dmg:      "dmg",
		&a.Cec:      "cec",
		&a.Dcec:     "dcec",
		&a.Capermg:  "capermg",
		&a.Mgperk:   "mgperk",
	}
}

func (a *Analysis) TableName() string {
	return "Analysis"
}

func Insert(a *Analysis) int64 {
	var dbs goydb.Goydb = a
	id, err := goydb.Insert(dbs)
	if err != nil {
		panic(err)
	}

	return id
}

func GetById(aId int) *Analysis {
	var a Analysis

	db := goydb.Conn()
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM Analysis WHERE id = ?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(aId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var fieldId int
		var fieldKey string
		var ph float64
		var phk float64
		var ec float64
		var php float64
		var eofph float64
		var k float64
		var dk float64
		var ca float64
		var dca float64
		var mg float64
		var dmg float64
		var cec int
		var dcec float64
		var capermg float64
		var mgperk float64
		err = rows.Scan(&id, &fieldId, &fieldKey, &ph, &phk, &ec, &php, &eofph, &k, &dk, &ca, &dca, &mg, &dmg, &cec, &dcec, &capermg, &mgperk)
		if err != nil {
			panic(err)
		}
		a = Analysis{id, fieldId, fieldKey, ph, phk, ec, php, eofph, k, dk, ca, dca, mg, dmg, cec, dcec, capermg, mgperk}
	}

	return &a
}
