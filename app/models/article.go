package models

import (
	"runtime"
	"time"
	"net/http"
	"context"
	"github.com/labstack/echo/v4"
	// "encoding/json"
	database "github.com/HendricksK/sacosbego/app/database"
	extensions "github.com/HendricksK/sacosbego/app/extensions"
)

type Article struct {
	Id				*int 		`json:"id"`
	Name 			*string 	`json:"name"` 
	Data 			*string 	`json:"data"`
	Uri 			*string 	`json:"uri"`
	Author			*string 	`json:"author"`
	Tags			*string		`json:"tags"`  
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}


type ArticleAggregate struct {
	ArticleId		*int 		`json:"article_id"`
	Tags 			*string 	`json:"tags"` 
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}

var article_model string = "article"
var article_model_aggregate string = "article_aggregate"


func GetArticle(id string) Article {

	var article Article
	
	db := database.Open()

	// CREATE CONN
	fields := []string{
		article_model + ".id",
		article_model + ".name",
		article_model + ".data",
		article_model + ".uri",
		article_model + ".author",
		article_model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, article_model, article_model_aggregate)
	
	err := db.QueryRow(selectQuery + " WHERE id =" + id).Scan(
			&article.Id, 
			&article.Name, 
			&article.Data, 
			&article.Uri, 
			&article.Author,
			&article.Tags)

	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
	}

	database.Close(db)
	
	// CLOSE CONN
	return article
}

func GetArticles() []Article {

	var articles []Article

	db := database.Open()

	// CREATE CONN
	fields := []string{
		article_model + ".id",
		article_model + ".name",
		article_model + ".data",
		article_model + ".uri",
		article_model + ".author",
		article_model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, article_model, article_model_aggregate)
	
	rows, err := db.Query(selectQuery)
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
		return articles
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
			_, filename, line, _ := runtime.Caller(1)
			extensions.Log(err.Error(), filename, line)
		    panic(err)
		}

		articles = append(articles, article)
	}

	database.Close(db)

	// CLOSE CONN
	return articles
}

func CreateArticle(c echo.Context) int {
	var article Article

	db := database.Open()

	err := c.Bind(&article); if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
	    return http.StatusBadRequest
	}

	fields := []string{
		"name",
		"data",
		"uri",
		"author"}

	insertQuery := BuildInsertQuery(fields, article_model)

	result, err := db.ExecContext(context.Background(), insertQuery, *article.Name, *article.Data, *article.Uri, *article.Author)
	if err != nil {
		panic(err)
		return http.StatusBadRequest
	}
	articleId, err := result.LastInsertId()
	
	aggregrateFields := []string{
		"article_id",
		"tags"}

	insertAggregateQuery := BuildInsertQuery(aggregrateFields, article_model_aggregate)

	_, err = db.ExecContext(context.Background(), insertAggregateQuery, articleId, *article.Tags) 
	if err != nil {	
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
		return http.StatusBadRequest
	}

	database.Close(db)

	return http.StatusCreated
}


// TODO: build out actual update
func UpdateArticle(c echo.Context) int {
	return http.StatusForbidden
}

// TODO: build out actual delete
// maybe add an extra column, is_ative?
func DeleteArticle(c echo.Context) int {
	return http.StatusForbidden
}