package conf

import (
	"Studs/pkg/models"
	"log"

	"github.com/BurntSushi/toml"
)

func FacebookClient() (string, string) {
	var Config models.Tomlconf
	if _, err := toml.DecodeFile("pkg/config/env/config.toml", &Config); err != nil {
		log.Fatal(err)
	}
	return Config.FacebookLogin.ClientID, Config.FacebookLogin.ClientSecret
}
