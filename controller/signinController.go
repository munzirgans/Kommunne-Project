package controller

import (
	"Studs/pkg/config"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var Store = sessions.NewCookieStore([]byte("MunzirStuds"))

//SigninController Function
func SigninController(w http.ResponseWriter, r *http.Request) {
	var passDB string
	var username string
	data := map[string]interface{}{
		"err": "Invalid username or password",
	}
	email := r.FormValue("your_email")
	password := r.FormValue("your_pass")
	errs := config.DB.QueryRow("select password,name from user where email=?", email).Scan(&passDB, &username)
	if errs != nil {
		tmp, err := template.ParseFiles("views/signin.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmp.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	errs = bcrypt.CompareHashAndPassword([]byte(passDB), []byte(password))
	if errs != nil {
		tmp, err := template.ParseFiles("views/signin.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = tmp.Execute(w, data)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
	session, _ := Store.Get(r, "session")
	session.Values["username"] = username
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DelSes(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

//TesSes Function
func TesSes(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		return
	}
	username, ok := untyped.(string)
	if !ok {
		return
	}
	fmt.Println(username)
}
