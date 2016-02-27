package dbconf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Db   string
	User string
	Pass string
}

func GetDBConfig() Config {
	var conf Config

	p, _ := os.Getwd()
	if strings.Index(p, "\\") > 0 {
		p = strings.Replace(p, "\\", "/", -1)
	}
	p = strings.Replace(p, "/dbconf", "", 1)

	jsonString, err := ioutil.ReadFile(p + "/conf/mysql.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonString, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
