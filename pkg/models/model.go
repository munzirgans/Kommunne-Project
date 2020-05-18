package models

//Tomlconf config model
type Tomlconf struct {
	Database struct {
		Name     string `toml:"name"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Host     string `toml:"host"`
		Port     string `toml:"port"`
	}
	GoogleLogin struct {
		ClientID     string `toml:"clientID"`
		ClientSecret string `toml:"clientSecret"`
	}
}

type GoogleProfile struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
