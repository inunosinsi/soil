package view

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	
	"../model/analysis"
)

type graphHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func NewGraphHandler(filename string) graphHandler {
	return graphHandler{filename: filename}
}

func (h *graphHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var ph float64
	var ec float64
	var eofph float64
	var k float64
	var ca float64
	var mg float64
	var cec float64
	
	/** @ToDo 塩基飽和度の計算がわかんね **/

	a := analysis.GetById(2)
	ph = a.Ph / 6 * 50
	ec = a.Ec / 0.5 * 50
	eofph = a.Eofph / 60 * 50
	k = a.K / 51 * 50
	ca = a.Ca / 153 * 50
	mg = a.Mg / 44 * 50
	cec = a.Cec / 15 * 50

	h.once.Do(func() {
		h.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				h.filename)))
	})
	data := map[string]interface{}{
		"Ph": ph,
		"Ec": ec,
		"Eofph": eofph,
		"K": k,
		"Ca": ca,
		"Mg": mg,
		"Cec": cec,
	}

	h.templ.Execute(w, data)
}
