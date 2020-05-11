package template

import (
	"Studs/controller"
	"net/http"
)

func IndexTemplate(w http.ResponseWriter, r *http.Request) {
	var tmpl = ParseFiles("index.html")
	var data = map[string]interface{}{
		"login": true,
	}
	session, _ := controller.Store.Get(r, "session")
	_, notok := session.Values["username"]
	if !notok {
		data["login"] = false
	}
	var err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
