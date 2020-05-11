package controller

import (
	"Studs/pkg/config"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//SignupController Function
func SignupController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("name")
		password := r.FormValue("pass")
		email := r.FormValue("email")
		hashedpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		_, errs := config.DB.Exec("insert into user (name,email,password) values (?,?,?)", username, email, hashedpass)
		if errs != nil {
			http.Error(w, errs.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
