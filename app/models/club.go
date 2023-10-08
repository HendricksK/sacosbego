package models

import (
	// "runtime"
	"time"
	// "net/http"
	// "context"
	// "github.com/labstack/echo/v4"
	// "encoding/json"
	// database "github.com/HendricksK/sacosbego/app/database"
	// extensions "github.com/HendricksK/sacosbego/app/extensions"
)

type Club struct {
	Id				*int 		`json:"id"`
	Name 			*string 	`json:"name"` 
	Data 			*string 	`json:"data"`
	Uri 			*string 	`json:"uri"`
	Author			*string 	`json:"author"`
	Tags			*string		`json:"tags"`  
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}


type ClubAggregate struct {
	ClubId			*int 		`json:"club_id"`
	Tags 			*string 	`json:"tags"` 
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}

// var article_model string = "article"
// var article_model_aggregate string = "article_aggregate"

// func GetClub (id int) {

// }