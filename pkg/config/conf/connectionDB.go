package conf

import (
	"Studs/pkg/models"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//ConnectionDB to make url connection db
func ConnectionDB() string {
	var config models.Tomlconf
	f, err := os.Open("pkg/config/env/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	toml.Unmarshal(content, &config)
	connURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	return connURL
}
