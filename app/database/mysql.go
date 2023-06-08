package mysql_connection

import (
	_ "github.com/go-sql-driver/mysql"
)

func Create() _ {
	db, err := sql.Open("mysql", "dev-sacos:CorsairTX850M@/sacos_dev")
	return db
}