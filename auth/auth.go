package auth 

import (
	"log"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

type Auth struct {
	Id		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Token	string 	`json:"token"`
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetSacosUploadToken(token_name string) Auth {

	var auth Auth
	err := db.QueryRow("SELECT * FROM auth WHERE name = $1;", token_name).Scan(&auth.Id, &auth.Name, &auth.Token)
	if err != nil {
	    log.Println(err)
	}

	return Auth {
		Id: auth.Id,
		Name: auth.Name,
		Token: auth.Token }
}
