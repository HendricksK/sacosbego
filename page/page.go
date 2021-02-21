package page

import (
	"log"
	"os"
	"database/sql"
	post "github.com/HendricksK/sacosbego/post"

	_ "github.com/lib/pq"
)

type Page struct {
	Id 			int 		`json:"id"`
	Page_data 	string 		`json:"page_data"`
	Datetime    string 		`json:"datatime"`
	Post_data	[]post.Post `json:"post_data"`
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetPage(page_id int) Page {
	var page Page
	err := db.QueryRow("SELECT * FROM pages WHERE id = $1;", page_id).Scan(&page.Id, &page.Page_data, &page.Datetime)
	// https://stackoverflow.com/questions/25501875/error-in-importing-custom-packages-in-go-lang
	page.Post_data = post.GetPosts(page_id)
	if err != nil {
	    log.Println(err)
	}

	return page
}
