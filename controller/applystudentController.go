package controller

import (
	"Studs/pkg/config"
	"fmt"
	"net/http"
)

func ApplyStudentController(w http.ResponseWriter, r *http.Request) {
	var userid int
	session, _ := Store.Get(r, "session")
	username, _ := session.Values["username"]
	username = username.(string)
	err := config.DB.QueryRow("select user_id from user where name = ?", username).Scan(&userid)
	if err != nil {
		fmt.Println(err)
		return
	}
	school := r.FormValue("school")
	goal := r.FormValue("goal")
	linkedin := r.FormValue("linkedin")
	_, err = config.DB.Exec("insert into studentapply (user_id,school,goal,linkedin) values (?,?,?,?)",
		userid,
		school,
		goal,
		linkedin,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
