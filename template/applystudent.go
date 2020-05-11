package template

import (
	"Studs/controller"
	"net/http"
	"path"
	"text/template"
)

func ApplystudentTemplate(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "std-form.html")
	session, _ := controller.Store.Get(r, "session")
	_, notok := session.Values["username"]
	if !notok {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}
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
