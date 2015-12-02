package webview

import (
	"html/template"
	"net/http"
)

type Page struct {
	Body string
	tpl  *template.Template
}

func LoadPage(path string) (*Page, error) {
	tpl, err := template.ParseFiles(path)

	return &Page{tpl: tpl}, err
}

func (p *Page) WritePage(w http.ResponseWriter, origins []string) error {
	return p.tpl.Execute(w, origins)
}
