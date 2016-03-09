package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../model/analysis"
)

type jsonHandler struct {
}

func NewJsonHandler() jsonHandler {
	return jsonHandler{}
}

func (h *jsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/**
	 * @ToDo field_keyで取得できる様にしたい
	 */
	a := analysis.GetById(2)

	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(out))
}
