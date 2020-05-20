package template

import (
	"Studs/controller"
	"net/http"
	"path"
	"text/template"
)

//SigninTemplate Temp
func SigninTemplate(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "signin.html")
	tmpl, err := template.ParseFiles(filepath)
	data := map[string]interface{}{
		"google":   controller.GoogleAccountUrlRedirect,
		"facebook": controller.FacebookAccountUrlRedirect,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
