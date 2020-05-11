package template

import (
	"html/template"
	"net/http"
)

func TesTemplate(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/part/head.html",
		"views/part/header.html",
		"views/tes.html",
		"views/part/script.html",
		"views/part/footer.html",
	))
	var err = tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
