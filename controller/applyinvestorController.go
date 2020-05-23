package controller

import (
	"Studs/pkg/config"
	"fmt"
	"net/http"
)

func ApplyInvestorController(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	username, _ := session.Values["username"]
	_, err := config.DB.Exec("update user set role = ? where name = ?",
		"investor",
		username.(string),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/marketplace", http.StatusSeeOther)
}
