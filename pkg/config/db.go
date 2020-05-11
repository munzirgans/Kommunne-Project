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
