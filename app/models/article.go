package models

import (
	// "log"
	"time"
	"net/http"
	"context"
	"github.com/labstack/echo/v4"
	// "encoding/json"
	database "github.com/HendricksK/sacosbego/app/database"

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

var model string = "article"
var model_aggregate string = "article_aggregate"


func GetArticle(id string) Article {

	var article Article
	
	db := database.Open()

	// CREATE CONN
	fields := []string{
		model + ".id",
		model + ".name",
		model + ".data",
		model + ".uri",
		model + ".author",
		model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, model, model_aggregate)
	
	err := db.QueryRow(selectQuery + " WHERE id =" + id).Scan(
			&article.Id, 
			&article.Name, 
			&article.Data, 
			&article.Uri, 
			&article.Author,
			&article.Tags)

	if err != nil {
		panic(err)
		return article
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
		model + ".id",
		model + ".name",
		model + ".data",
		model + ".uri",
		model + ".author",
		model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, model, model_aggregate)
	
	rows, err := db.Query(selectQuery)
	if err != nil {
		panic(err)
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
		panic(err)
	    return http.StatusBadRequest
	}

	fields := []string{
		"name",
		"data",
		"uri",
		"author"}

	insertQuery := BuildInsertQuery(fields, model)

	result, err := db.ExecContext(context.Background(), insertQuery, *article.Name, *article.Data, *article.Uri, *article.Author)
	if err != nil {
		panic(err)
		return http.StatusBadRequest
	}
	articleId, err := result.LastInsertId()
	
	aggregrateFields := []string{
		"article_id",
		"tags"}

	insertAggregateQuery := BuildInsertQuery(aggregrateFields, model_aggregate)

	_, err = db.ExecContext(context.Background(), insertAggregateQuery, articleId, *article.Tags) 
	if err != nil {	
		panic(err)
		return http.StatusBadRequest
	}

	database.Close(db)

	return http.StatusCreated
}

// func CreateArticleAggregate(c echo.Context) int {
// 	var articleAggregate ArticleAggregate

// 	db := database.Open()

// 	err := c.Bind(&articleAggregate); if err != nil {
// 		panic(err)
// 	    return http.StatusBadRequest
// 	}

// 	aggregrateFields := []string{
// 		"article_id",
// 		"tags"}

// 	insertAggregateQuery := BuildInsertQuery(aggregrateFields, model_aggregate)

// 	_, err = db.ExecContext(context.Background(), insertAggregateQuery, articleAggregate.ArticleId, *articleAggregate.Tags) 
// 	if err != nil {	
// 		panic(err)
// 		return http.StatusBadRequest
// 	}

// 	database.Close(db)

// 	return http.StatusCreated
// }

func UpdateArticle(c echo.Context) int {
	return http.StatusForbidden
}

func DeleteArticle(c echo.Context) int {
	return http.StatusForbidden
}