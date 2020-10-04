package page

import (
	"log"
	"os"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Page struct {
	Id				int
	Page_data 		string
	Post_data		Post    
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func getPage(page_id int) Page {
	var page Page
	err := db.QueryRow("SELECT * FROM page WHERE id = $1;", page_id).Scan(&page.Id, &page.Page_data)
	page.Post_data := post.getPost(page_id)
	if err != nil {
	    log.Println(err)
	}

	fmt.Print(page)

	return page
}
