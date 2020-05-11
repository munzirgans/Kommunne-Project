package models

//Tomlconf config model
type Tomlconf struct {
	Database struct {
		Name     string
		User     string
		Password string
		Host     string
		Port     string
	}
}
