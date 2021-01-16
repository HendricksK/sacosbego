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
	Rider_data 		string 
	Url 			string 
	Datetime 		string
	Updatetime 		string
}

type RiderId struct {
	Id				int
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetRider(rider_id int) Rider {

	var rider Rider
	err := db.QueryRow("SELECT * FROM rider WHERE id = $1;", rider_id).Scan(&rider.Id, &rider.Name, &rider.Rider_data, &rider.Url, &rider.Datetime, &rider.Updatetime)
	if err != nil {
	    log.Println(err)
	}

	return Rider {
		Id: rider.Id,
		Name: rider.Name,
		Rider_data: rider.Rider_data,
		Url: rider.Url,
		Datetime: rider.Datetime,
		Updatetime: rider.Updatetime,
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

