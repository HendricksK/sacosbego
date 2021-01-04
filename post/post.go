package post

import (
	"log"
	"os"
	"database/sql"
	// "encoding/json"

	_ "github.com/lib/pq"
)

type Post struct {
	Id 			int
	Page_id 	int
	Section 	string
	Page_data 	[]uint8
	Page_data_string string
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetPosts(page_id int) []Post {
	var posts []Post

	rows, err := db.Query("SELECT id, page_id, section, page_data FROM posts WHERE page_id = $1 ORDER BY id ASC;", page_id)
	if err != nil {
	    log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		
		var post Post
		err = rows.Scan(&post.Id, &post.Page_id, &post.Section, &post.Page_data)

		if err != nil {
		    log.Println(err)
		}
		post.Page_data_string = string(post.Page_data)
		posts = append(posts, post)	
	}

	return posts
}

func GetPostSection(page_id int, section string) Post {
	var post Post

	err := db.QueryRow("SELECT id, page_id, section, page_data FROM posts WHERE page_id = $1 AND section = $2;", page_id, section).Scan(&post.Id, &post.Page_id, &post.Section, &post.Page_data)
	if err != nil {
	    log.Println(err)
	}

	post.Page_data_string = string(post.Page_data)

	return post
}
