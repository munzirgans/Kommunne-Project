package template

import (
	"Studs/controller"
	"net/http"
)

func MarketplaceTemplate(w http.ResponseWriter, r *http.Request) {
	var tmpl = ParseFiles("marketplace.html")
	session, _ := controller.Store.Get(r, "session")
	username, notok := session.Values["username"]
	var data = map[string]interface{}{
		"login": true,
		"name":  username.(string),
	}
	if !notok {
		data["login"] = false
	}
	var err = tmpl.ExecuteTemplate(w, "marketplace", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
