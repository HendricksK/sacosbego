package models

import (
	"runtime"
	"fmt"
	"time"
	"net/http"
	// "strconv"
	// "context"
	"github.com/labstack/echo/v4"
	// "encoding/json"
	database "github.com/HendricksK/sacosbego/app/database"
	extensions "github.com/HendricksK/sacosbego/app/extensions"
)

type Rider struct {
	Id				*int 		`json:"id"`
	FirstName 		*string 	`json:"first_name"`
	LastName 		*string 	`json:"last_name"`
	DateOfBirth		*time.Time	`json:"date_of_birth"` // Time will always be 00:00:00 we just want a time object
	Data 			*string 	`json:"data"`
	Uri 			*string 	`json:"uri"`
	Tags			*string		`json:"tags"`
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
	Images			[]Image	`json:"images"`
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

var rider_model string = "rider"
var rider_model_aggregate string = "rider_aggregate"


func GetRider(id string) Rider {

	var rider Rider
	
	db := database.Open()

	// CREATE CONN
	fields := []string{
		rider_model + ".id",
		rider_model + ".first_name",
		rider_model + ".last_name",
		rider_model + ".date_of_birth",
		rider_model + ".data",
		rider_model + ".uri",
		rider_model_aggregate + ".tags"}

	var selectQuery = BuildSelectQueryWithAggregate(fields, rider_model, rider_model_aggregate)
	
	err := db.QueryRow(selectQuery + " WHERE id = " + id).Scan(
			&rider.Id, 
			&rider.FirstName, 
			&rider.LastName,
			&rider.DateOfBirth,  
			&rider.Data, 
			&rider.Uri,
			&rider.Tags)

	rider.Images = GetImages(rider_model, id)

	fmt.Printf("%v", rider.Images)

	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		extensions.Log(err.Error(), filename, line)
	}

	database.Close(db)
	
	// CLOSE CONN
	return rider
}

// TODO: build out actual update
func GetRiders() []Rider {

	var riders []Rider

	db := database.Open()

	database.Close(db)

	return riders
}

// TODO: build out actual update
func CreateRider(c echo.Context) int {
	return http.StatusForbidden
}

// TODO: build out actual update
func UpdateRider(c echo.Context) int {
	return http.StatusForbidden
}

// TODO: build out actual delete
// maybe add an extra column, is_ative?
func DeleteRider(c echo.Context) int {
	return http.StatusForbidden
}