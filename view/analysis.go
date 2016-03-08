package view

import (
	"net/http"
	"sync"
	"text/template"
)

type analysisHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func NewAnalysisHandler(filename string) analysisHandler {
	return analysisHandler{filename: filename}
}

func (h *analysisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
