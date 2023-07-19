package models

import (
	"runtime"
	"time"
	// "net/http"
	// "context"
	// "github.com/labstack/echo/v4"
	// "encoding/json"
	database "github.com/HendricksK/sacosbego/app/database"
	extensions "github.com/HendricksK/sacosbego/app/extensions"
)

type Post struct {
	Id				*int 		`json:"id"`
	PageId 			*int 		`json:"page_id"`
	Section			*string 	`json:"section"`
	Data 			*string 	`json:"data"`
	DataJson		*string 	`json:"data_json"`
	Order 			*int 		`json:"order"`
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}

var post_model = "post"


func GetPostsViaPageId(page_id string) []Post {
	var posts []Post
	
	db := database.Open()

	// CREATE CONN
	fields := []string{
		post_model + ".id",
		post_model + ".page_id",
		post_model + ".section",
		post_model + ".data",
		post_model + ".data_json",
		post_model + ".order",
	}

	var selectQuery = BuildSelectQuery(fields, post_model)
	
	rows, err := db.Query(selectQuery + " WHERE page_id = " + page_id)
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		
		err = rows.Scan(
			&post.Id, 
			&post.PageId, 
			&post.Section, 
			&post.Data,  
			&post.DataJson,
			&post.Order)

		if err != nil {
			_, filename, line, _ := runtime.Caller(1)
			extensions.Log(err.Error(), filename, line)
			panic(err)
		}

		posts = append(posts, post)
	}

	database.Close(db)
	
	return posts
}