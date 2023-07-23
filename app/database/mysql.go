package database

import (
	// "log"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// 192.168.0.134:3309
var db_host = "172.17.0.3:3306";

func Open() *sql.DB {
	// parseTime=true https://github.com/go-sql-driver/mysql#timetime-support
	database, err := sql.Open("mysql", "dev-sacos:CorsairTX850M@tcp(" + db_host + ")/sacos_dev?parseTime=true")
	// id:password@tcp(your-amazonaws-uri.com:3306)/dbname
	if err != nil {
		panic(err)
		
	}
	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	return database
}

func Close(conn *sql.DB) {
	conn.Close()
}
