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
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetArticle(article_id int) Article {

	var article Article
	err := db.QueryRow("SELECT * FROM article WHERE id = $1", article_id).Scan(&article.Id, &article.Name, &article.Article_data, &article.Url, &article.Datetime)
	
	if err != nil {
	    log.Println(err)
	}

	return Article{
		Id: article.Id,
		Name: article.Name,
		Article_data: article.Article_data,
		Url: article.Url,
		Datetime: article.Datetime}
}

