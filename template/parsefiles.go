package template

import (
	"html/template"
)

func ParseFiles(filepath string) *template.Template {
	result := template.Must(template.ParseFiles(
		"views/"+filepath,
		"views/part/head.html",
		"views/part/header.html",
		"views/part/footer.html",
		"views/part/script.html",
	))
	return result
}
