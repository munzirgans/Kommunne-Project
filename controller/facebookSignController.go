package controller

import (
	"Studs/pkg/config"
	"Studs/pkg/config/conf"
	"Studs/pkg/models"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var FacebookOauthConfig *oauth2.Config
var FacebookAccountUrlRedirect string

func SetFOauthConfig() *oauth2.Config {
	clientID, clientSecret := conf.FacebookClient()
	foauthconfig := &oauth2.Config{
		RedirectURL:  "http://kommunne.herokuapp.com/facebooksign",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"email", "public_profile", "user_photos"},
		Endpoint:     facebook.Endpoint,
	}
	return foauthconfig
}

func SetFAccountUrlRedirect() string {
	url := FacebookOauthConfig.AuthCodeURL("kommunnestate")
	return url
}

func FacebookSigninController(w http.ResponseWriter, r *http.Request) {
	var fprof models.FacebookProfile
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var multiplerow bool
	if r.FormValue("state") != "kommunnestate" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	token, err := FacebookOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := http.Get("https://graph.facebook.com/me?fields=name,email&access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal([]byte(string(content)), &fprof); err != nil {
		fmt.Println(err)
		return
	}
	_ = config.DB.QueryRow("select if(count(*) > 0, 'true','false') from user where email = ?", fprof.Email).Scan(&multiplerow)
	if multiplerow {
		session, _ := Store.Get(r, "session")
		session.Values["username"] = fprof.Name
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	_, errs := config.DB.Exec("insert into user(name,password,email) values (?,?,?)",
		fprof.Name,
		"facebooklogin",
		fprof.Email,
	)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
	session, _ := Store.Get(r, "session")
	session.Values["username"] = fprof.Name
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
