package models

import (
	// "fmt"
	"runtime"
	"time"
	"net/http"
	// "context"
	"github.com/labstack/echo/v4"
	// "encoding/json"
	database "github.com/HendricksK/sacosbego/app/database"
	extensions "github.com/HendricksK/sacosbego/app/extensions"
)

type Page struct {
	Id				*int 		`json:"id"`
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
	Posts			[]Post		`json:"posts"`
}


var page_model = "page"
var page_foreign_join = "post"


func GetPage(id string) Page {

	var page Page
	
	db := database.Open()

	// CREATE CONN
	fields := []string{
		page_model + ".id",
		page_model + ".created_at",
		page_model + ".updated_at"}

	var selectQuery = BuildSelectQuery(fields, page_model)
	
	err := db.QueryRow(selectQuery + " WHERE id =" + id ).Scan(
		&page.Id,
		&page.CreatedAt, 
		&page.UpdatedAt)

	page.Posts = GetPostsViaPageId(id)

	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
	}

	database.Close(db)
	
	// CLOSE CONN
	return page
}

func GetPages() []Page {
	var pages []Page
	return pages
}

func CreatePage(c echo.Context) int {
	return http.StatusForbidden
}


// TODO: build out actual update
func UpdatePage(c echo.Context) int {
	return http.StatusForbidden
}

// TODO: build out actual delete
// maybe add an extra column, is_ative?
func DeletePage(c echo.Context) int {
	return http.StatusForbidden
}