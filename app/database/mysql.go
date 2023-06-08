package database

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Create() *sql.DB {
	DB, err := sql.Open("mysql", "dev-sacos:CorsairTX850M@/sacos_dev")
	if err != nil {
		log.Println(err)
		// need to throw an error here
	}
	return DB
}

func Close(conn *sql.DB) {
	conn.Close()
}
