package models

import (
	"log"
	// "os"
	// "database/sql"
	database "github.com/HendricksK/sacosbego/app/database"

	// _ "github.com/lib/pq"

)

type Article struct {
	Id				int 	`json:"id"`
	Name 			string 	`json:"name"` 
	Article_data 	string 	`json:"article_data"`
	Url 			string 	`json:"url"`
	Datetime 		string 	`json:"datetime"`
	Author			string 	`json:"author"`
}

func GetArticle() string {
	log.Println("I am an article")
	db := database.Create()
	database.Close(db)
	return "move it to the left... yeah"
}

// type ArticleId struct {
// 	Id				int 	`json:"id"`
// }

// var connStr = os.Getenv("DATABASE_URL")
// var db, err = sql.Open("postgres", connStr)

// func GetArticle(article_id int) Article {

// 	var article Article
// 	err := db.QueryRow("SELECT * FROM article WHERE id = $1;", article_id).Scan(&article.Id, &article.Name, &article.Article_data, &article.Url, &article.Datetime, &article.Author)
// 	if err != nil {
// 	    log.Println(err)
// 	}

// 	return Article {
// 		Id: article.Id,
// 		Name: article.Name,
// 		Article_data: article.Article_data,
// 		Url: article.Url,
// 		Datetime: article.Datetime,
// 		Author: article.Author }
// }

// func GetArticles() []Article {

// 	var articles []Article
// 	rows, err := db.Query("SELECT id, name, datetime, author FROM article")
// 	if err != nil {
// 	    log.Println(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
		
// 		var article Article
// 		err = rows.Scan(&article.Id, &article.Name, &article.Datetime, &article.Author)
// 		if err != nil {
// 		    log.Println(err)
// 		}

// 		articles = append(articles, article)	
// 	}

// 	return articles
	
// }

// func GetArticlesIds() []ArticleId{
	
// 	var articles []ArticleId
// 	rows, err := db.Query("SELECT id FROM article")
// 	if err != nil {
// 	    log.Println(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
		
// 		var article ArticleId
// 		err = rows.Scan(&article.Id)
// 		if err != nil {
// 		    log.Println(err)
// 		}

// 		articles = append(articles, article)	
// 	}

// 	return articles
// }

