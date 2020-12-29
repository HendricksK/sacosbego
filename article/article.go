package article

import (
	"log"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

type Article struct {
	Id				int
	Name 			string 
	Article_data 	string 
	Url 			string 
	Datetime 		string
	Author			string
}

type ArticleId struct {
	Id				int
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetArticle(article_id int) Article {

	var article Article
	err := db.QueryRow("SELECT * FROM article WHERE id = $1;", article_id).Scan(&article.Id, &article.Name, &article.Article_data, &article.Url, &article.Datetime, &article.Author)
	if err != nil {
	    log.Println(err)
	}

	return Article {
		Id: article.Id,
		Name: article.Name,
		Article_data: article.Article_data,
		Url: article.Url,
		Datetime: article.Datetime,
		Author: article.Author }
}

func GetArticles() []Article {

	log.Println(connStr)

	var articles []Article
	rows, err := db.Query("SELECT id, name, datetime, author FROM article")
	if err != nil {
	    log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		
		var article Article
		err = rows.Scan(&article.Id, &article.Name, &article.Datetime, &article.Author)
		if err != nil {
		    log.Println(err)
		}

		articles = append(articles, article)	
	}

	return articles
	
}

func GetArticlesIds() []ArticleId{
	var articles []ArticleId
	rows, err := db.Query("SELECT id FROM article")
	if err != nil {
	    log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		
		var article ArticleId
		err = rows.Scan(&article.Id)
		if err != nil {
		    log.Println(err)
		}

		articles = append(articles, article)	
	}

	return articles
}

