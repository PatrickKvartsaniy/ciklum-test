package tools

import (
	"html/template"
	"net/http"
)

// RenderPage render  template or error page
func RenderPage(w http.ResponseWriter, e string) {
	tmpl, err := template.ParseFiles("index.html")
	CheckErr(err)
	tmpl.Execute(w, e)
}
