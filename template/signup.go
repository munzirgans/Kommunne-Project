package template

import (
	"net/http"
	"path"
	"text/template"
)

//SignupTemplate Template
func SignupTemplate(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "signup.html")
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
