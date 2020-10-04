package post

import (
	"log"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

type Post struct {
	Id 			int
	Page_id 	int
	Section 	string
	Page_data 	string
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetPosts(post_id int) []Post {
	var posts []Post

	rows, err := db.Query("SELECT id, page_id, sectoin, page_data FROM posts WHERE page_id = $1;", post_id)
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

		posts = append(posts, post)	
	}

	return posts
}
