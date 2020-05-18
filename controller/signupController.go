package controller

import (
	"Studs/pkg/config"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

//SignupController Function
func SignupController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var errmsg = false
		data := map[string]interface{}{
			"username":    false,
			"email":       false,
			"password":    false,
			"confirmpass": false,
		}
		username := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("pass")
		confpass := r.FormValue("re_pass")
		reemail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !nameValid(username) {
			data["username"] = true
			errmsg = true
		} else if !reemail.MatchString(email) {
			data["email"] = true
			errmsg = true
		} else if !isValid(password) {
			data["password"] = true
			errmsg = true
		} else if confpass != password {
			data["confirmpass"] = true
			errmsg = true
		}
		if errmsg {
			tmp, err := template.ParseFiles("views/signup.html")
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

func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func IsDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isValid(s string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(s) >= 8 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}

func nameValid(s string) bool {
	letter := false
	digit := false
	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			letter = true
		case !unicode.IsDigit(r):
			digit = true
		}
	}
	return letter && digit

}
