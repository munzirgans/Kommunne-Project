package template

import (
	"Studs/controller"
	"Studs/pkg/config"
	"fmt"
	"net/http"
)

func MarketplaceTemplate(w http.ResponseWriter, r *http.Request) {
	var roleinv bool
	var tmpl = ParseFiles("marketplace.html")
	session, _ := controller.Store.Get(r, "session")
	username, notok := session.Values["username"]
	if !notok {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	errs := config.DB.QueryRow("select if(role=?,?,?) from user where name = ?",
		"investor",
		"true",
		"false",
		username,
	).Scan(&roleinv)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	if !roleinv {
		http.Redirect(w, r, "/apply-investor", http.StatusSeeOther)
		return
	}
	var data = map[string]interface{}{
		"login": true,
		"name":  username.(string),
	}
	var err = tmpl.ExecuteTemplate(w, "marketplace", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
