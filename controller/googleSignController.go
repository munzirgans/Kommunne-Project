package controller

import (
	"Studs/pkg/config"
	"Studs/pkg/config/conf"
	"Studs/pkg/models"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig *oauth2.Config
var GoogleAccountUrlRedirect string

func SetGOauthConfig() *oauth2.Config {
	clientID, clientSecret := conf.GoogleClient()
	GOauthConfig := &oauth2.Config{
		RedirectURL:  "http://kommunne.herokuapp.com/googlesign",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return GOauthConfig
}

func SetGAccountUrlRedirect() string {
	url := GoogleOauthConfig.AuthCodeURL("studsstate")
	return url
}

func GoogleSigninController(w http.ResponseWriter, r *http.Request) {
	var gprof models.GoogleProfile
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var multiplerow bool
	if r.FormValue("state") != "studsstate" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	token, err := GoogleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal([]byte(string(content)), &gprof); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = config.DB.QueryRow("select if(count(*) > 0, 'true','false') from user where email = ?", gprof.Email).Scan(&multiplerow)
	if multiplerow {
		session, _ := Store.Get(r, "session")
		session.Values["username"] = gprof.Name
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	_, errs := config.DB.Exec("insert into user(name,password,email) values (?,?,?)",
		gprof.Name,
		"googlelogin",
		gprof.Email,
	)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
	session, _ := Store.Get(r, "session")
	session.Values["username"] = gprof.Name
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
