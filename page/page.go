package page

import (
	"log"
	"os"
	"database/sql"
	"fmt"
	"github.com/HendricksK/sacosbego/post"

	_ "github.com/lib/pq"
)

type Page struct {
	Id 			int
	Page_data 	string
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetPage(page_id int) Page {
	var page Page
	err := db.QueryRow("SELECT * FROM page WHERE id = $1;", page_id).Scan(&page.Id, &page.Page_data)
	// https://stackoverflow.com/questions/25501875/error-in-importing-custom-packages-in-go-lang
	post.GetPosts(page_id)
	// page.Post_data := post.getPost(page_id)
	if err != nil {
	    log.Println(err)
	}

	fmt.Print(page)

	return page
}
