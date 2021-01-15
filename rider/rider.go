package rider

import (
	"log"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

type Rider struct {
	Id				int
	Name 			string 
	rider_data 	string 
	Url 			string 
	Datetime 		string
	Author			string
}

type RiderId struct {
	Id				int
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func Getrider(rider_id int) Rider {

	var rider Rider
	err := db.QueryRow("SELECT * FROM rider WHERE id = $1;", rider_id).Scan(&rider.Id, &rider.Name, &rider.rider_data, &rider.Url, &rider.Datetime, &rider.Author)
	if err != nil {
	    log.Println(err)
	}

	return Rider {
		Id: rider.Id,
		Name: rider.Name,
		rider_data: rider.rider_data,
		Url: rider.Url,
		Datetime: rider.Datetime,
		Author: rider.Author }
}

func Getriders() []Rider {

	var riders []Rider
	rows, err := db.Query("SELECT id, name, datetime, author FROM rider")
	if err != nil {
	    log.Println(err)
	}

	log.Println(rows)
	
	defer rows.Close()

	for rows.Next() {
		
		var rider Rider
		err = rows.Scan(&rider.Id, &rider.Name, &rider.Datetime, &rider.Author)
		if err != nil {
		    log.Println(err)
		}

		riders = append(riders, rider)	
	}

	return riders
	
}

func GetridersIds() []RiderId{
	
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

