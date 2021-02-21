package rider

import (
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

type Rider struct {
	Id				int    	`json:"id"`
	Name 			string 	`json:"name"` 
	Rider_data 		string 	`json:"rider_data"`
	Url 			string 	`json:"url"`
	Datetime 		string 	`json:"datetime"`
	Updatetime 		string 	`json:"updatetime"`
	Images 			[]Image `json:"images"`
}

type RiderId struct {
	Id				int 	`json:"id"`
}


type Image struct {
	Url 			string 	`json:"url"`
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetRider(rider_id int) Rider {

	var rider Rider
	err := db.QueryRow("SELECT * FROM rider WHERE id = $1;", rider_id).Scan(&rider.Id, &rider.Name, &rider.Rider_data, &rider.Url, &rider.Datetime, &rider.Updatetime)
	if err != nil {
	    log.Println(err)
	}

	var images []Image

	rows, err := db.Query("SELECT url FROM images WHERE info = $1 AND type = $2", rider_id, "rider")
	if err != nil {
	    log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {

		var image Image
		err = rows.Scan(&image.Url)
		if err != nil {
		    log.Println(err)
		}

		images = append(images, image)	
	}

	return Rider {
		Id: rider.Id,
		Name: rider.Name,
		Rider_data: rider.Rider_data,
		Url: rider.Url,
		Datetime: rider.Datetime,
		Updatetime: rider.Updatetime,
		Images: images,
	}
}

func GetRiders() []Rider {

	var riders []Rider
	rows, err := db.Query("SELECT id, name, Rider_data FROM rider")
	if err != nil {
	    log.Println(err)
	}

	log.Println(rows)
	
	defer rows.Close()

	for rows.Next() {
		
		var rider Rider
		err = rows.Scan(&rider.Id, &rider.Name, &rider.Rider_data)
		if err != nil {
		    log.Println(err)
		}

		riders = append(riders, rider)	
	}

	return riders
	
}

func GetRiderIds() []RiderId{
	
	var riders []RiderId
	rows, err := db.Query("SELECT id FROM rider")
	if err != nil {
	    log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		
		var rider RiderId
		err = rows.Scan(&rider.Id)
		if err != nil {
		    log.Println(err)
		}

		riders = append(riders, rider)	
	}

	return riders
}

