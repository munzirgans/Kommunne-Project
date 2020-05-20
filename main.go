package main

import (
	"Studs/controller"
	"Studs/pkg/config"
	"Studs/pkg/config/conf"
	render "Studs/template"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	connURL := conf.ConnectionDB()
	config.DB = config.DBConnect(connURL)
	controller.GoogleOauthConfig = controller.SetGOauthConfig()
	controller.GoogleAccountUrlRedirect = controller.SetGAccountUrlRedirect()
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("views/assets"))))
	// POST HANDLER
	r.HandleFunc("/signin", controller.SigninController).Methods("POST")
	r.HandleFunc("/signup", controller.SignupController).Methods("POST")
	r.HandleFunc("/googlesign", controller.GoogleSigninController).Methods("GET")
	r.HandleFunc("/iglogin", controller.InstagramController).Methods("POST")
	// GET HANDLER
	r.HandleFunc("/signin", render.SigninTemplate).Methods("GET")
	r.HandleFunc("/signup", render.SignupTemplate).Methods("GET")
	r.HandleFunc("/", render.IndexTemplate).Methods("GET")
	r.HandleFunc("/marketplace", render.MarketplaceTemplate).Methods("GET")
	r.HandleFunc("/ses", controller.TesSes).Methods("GET")
	r.HandleFunc("/apply-student", render.ApplystudentTemplate).Methods("GET")
	r.HandleFunc("/apply-investor", render.ApplyinvestorTemplate).Methods("GET")
	r.HandleFunc("/delses", controller.DelSes).Methods("GET")
	r.HandleFunc("/apply-investor", render.ApplyinvestorTemplate).Methods("GET")
	r.HandleFunc("/testemp", render.TesTemplate).Methods("GET")
	// r.HandleFunc("/instagram", render.InstagramSigninTemplate).Methods("GET")
	// fmt.Println("Connected to port 1234")
	// log.Fatal(http.ListenAndServe(":1234", r))
	fmt.Println("Connected to port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
