package tools

import (
	"ciklum/writer/tools"
	"html/template"
	"net/http"
)

// RenderPage render  template or error page
func RenderPage(w http.ResponseWriter, e string) {
	tmpl, err := template.ParseFiles("index.html")
	tools.CheckErr(err)
	tmpl.Execute(w, e)
}
