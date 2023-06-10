package models

import (
	"log"
	"time"
	// "fmt"
	// "strings"
	// "os"
	// "database/sql"
	database "github.com/HendricksK/sacosbego/app/database"

	// _ "github.com/lib/pq"

)

type Article struct {
	Id				int 	`json:"id"`
	Name 			*string 	`json:"name"` 
	Data 			*string 	`json:"article_data"`
	Uri 			*string 	`json:"uri"`
	Author			*string 	`json:"author"`
	Tags			*string  `json:"tags"`  
	CreatedAt 		*time.Time   	`json:"created_at"`
	UpdatedAt 		*time.Time     	`json:"updated_at"`
}

// CREATE TABLE article (
//     id int UNSIGNED NOT NULL AUTO_INCREMENT,
//     name varchar(255) NOT NULL,
//     data text NULL,
//     uri varchar(510) NULL,
//     author varchar(255) NULL,
//     created_at timestamp NOT NULL DEFAULT NOW(),
//     updated_at timestamp NULL ON UPDATE NOW(),
//     PRIMARY KEY (id)
// );
// CREATE TABLE article_aggregate (
// 	article_id int UNSIGNED NOT NULL AUTO_INCREMENT,
// 	tags JSON NULL,
// 	created_at timestamp NOT NULL DEFAULT NOW(),
//     updated_at timestamp NULL ON UPDATE NOW(),
// 	PRIMARY KEY (article_id)
// );

var db = database.Create()
var model string = "article"
var model_aggregate string = "article_aggregate"


func GetArticle(id int) Article {
	var article Article
	// CREATE CONN
	fields := []string{
		"id",
		"name",
		"data",
		"uri",
		"author"}

	var selectQuery = BuildSelectQuery(fields, model, id)
	
	err := db.QueryRow(selectQuery).Scan(
			&article.Id, 
			&article.Name, 
			&article.Data, 
			&article.Uri, 
			&article.Author)

	if err != nil {
		log.Println(err)
	}
	
	// CLOSE CONN
	return article
}

func GetArticles() []Article {
	var articles []Article
	// CREATE CONN
	fields := []string{
		model + ".id",
		model + ".name",
		model + ".data",
		model + ".uri",
		model + ".author",
		model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, model, model_aggregate)
	
	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var article Article
		
		err = rows.Scan(
			&article.Id, 
			&article.Name, 
			&article.Data, 
			&article.Uri, 
			&article.Author, 
			&article.Tags)

		if err != nil {
		    log.Println(err)
		}

		log.Println(article)
		articles = append(articles, article)
	}
	// CLOSE CONN
	return articles
}
