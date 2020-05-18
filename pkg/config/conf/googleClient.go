package conf

import (
	"Studs/pkg/models"
	"log"

	"github.com/BurntSushi/toml"
)

func GoogleClient() (string, string) {
	var Config models.Tomlconf
	if _, err := toml.DecodeFile("pkg/config/env/config.toml", &Config); err != nil {
		log.Fatal(err)
	}
	return Config.GoogleLogin.ClientID, Config.GoogleLogin.ClientSecret
}
