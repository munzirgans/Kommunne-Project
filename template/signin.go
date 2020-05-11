package template

import (
	"net/http"
	"path"
	"text/template"
)

//SigninTemplate Temp
func SigninTemplate(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "signin.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
