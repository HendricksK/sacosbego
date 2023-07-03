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

type Rider struct {
	Id				*int 		`json:"id"`
	FirstName 		*string 	`json:"first_name"`
	LastName 		*string 	`json:"last_name"`
	DateOfBirth		*time.Time	`json:"date_of_birth"` // Time will always be 00:00:00 we just want a time object
	Data 			*string 	`json:"data"`
	Uri 			*string 	`json:"uri"`
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
	Images			*[]Image		`json:"images"`
}


type RiderAggregate struct {
	RiderId			*int 		`json:"rider_id"`
	Tags 			*string 	`json:"tags"` 
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}

type RiderAccolade struct {
	RiderId			*int 		`json:"rider_id"`
	Name 			*string		`json:"name"`
	Year			*int 		`json:"year"`
	Tags 			*string 	`json:"tags"` 
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}

var model string = "rider"
var model_aggregate string = "rider_aggregate"


func GetRider(id string) Rider {

	var rider Rider
	
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