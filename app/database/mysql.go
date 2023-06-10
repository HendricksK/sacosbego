package database

import (
	"log"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func Create() *sql.DB {
	// parseTime=true https://github.com/go-sql-driver/mysql#timetime-support
	database, err := sql.Open("mysql", "dev-sacos:CorsairTX850M@tcp(192.168.0.134:3309)/sacos_dev?parseTime=true")
	// id:password@tcp(your-amazonaws-uri.com:3306)/dbname
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	if err != nil {
		log.Println(err)
		// need to throw an error here
	}

	return database
}

func Close(conn *sql.DB) {
	conn.Close()
}
