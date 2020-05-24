package template

import (
	"Studs/controller"
	"Studs/pkg/config"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type StudentApply struct {
	Name     string
	School   string
	Goal     int
	Linkedin string
}

func MarketplaceTemplate(w http.ResponseWriter, r *http.Request) {
	var sapply StudentApply
	var arr_sapply []StudentApply
	var roleinv bool
	session, _ := controller.Store.Get(r, "session")
	username, notok := session.Values["username"]
	funcs := template.FuncMap{"add": add, "percentage": percentage}
	tmpl := template.Must(template.New("marketplace").Funcs(funcs).ParseFiles("views/marketplace.html",
		"views/part/head.html",
		"views/part/header.html",
		"views/part/footer.html",
		"views/part/script.html",
	))
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
	rows, errs := config.DB.Query("select user.name, studentapply.school, studentapply.goal, studentapply.linkedin from studentapply inner join user on studentapply.user_id = user.user_id")
	if errs != nil {
		fmt.Println(errs)
		return
	}
	for rows.Next() {
		if err := rows.Scan(&sapply.Name, &sapply.School, &sapply.Goal, &sapply.Linkedin); err != nil {
			log.Fatal(err)
		}
		arr_sapply = append(arr_sapply, sapply)
	}
	var data = map[string]interface{}{
		"login": true,
		"name":  username.(string),
		"items": arr_sapply,
	}
	var err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func add(x, y int) int {
	return x + y
}

func percentage(x int) int {
	res := (x * 100) / 50000
	return res
}
