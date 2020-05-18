package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConnect(conn string) *sql.DB {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// [database]
// host = "us-cdbr-east-06.cleardb.net"
// port = "3306"
// user = "b2f07d608a35af"
// password = "f94be289"
// name = "heroku_73f670b93e27f80"

// [database]
// host = "localhost"
// port = "3306"
// user = "munzir"
// password = "munzirdev"
// name = "studs"
