package auth 

import (
	"log"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

type Ingest struct {
	Id				int 	`json:"id"`
	Form_upload_id 	string 	`json:"form_upload_id"`
	Form_data		json 	`json:"form_data"`
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func PostUploadData(formdata) Auth {
}
