package track

import (
	"log"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

type Track struct {
	Id				int
	Name 			string 
	track_data 	string 
	Url 			string 
	Datetime 		string
	Author			string
}

type TrackId struct {
	Id				int
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func Gettrack(track_id int) Track {

	var track Track
	err := db.QueryRow("SELECT * FROM track WHERE id = $1;", track_id).Scan(&track.Id, &track.Name, &track.track_data, &track.Url, &track.Datetime, &track.Author)
	if err != nil {
	    log.Println(err)
	}

	return Track {
		Id: track.Id,
		Name: track.Name,
		track_data: track.track_data,
		Url: track.Url,
		Datetime: track.Datetime,
		Author: track.Author }
}

func Gettracks() []Track {

	var tracks []Track
	rows, err := db.Query("SELECT id, name, datetime, author FROM track")
	if err != nil {
	    log.Println(err)
	}

	log.Println(rows)
	
	defer rows.Close()

	for rows.Next() {
		
		var track Track
		err = rows.Scan(&track.Id, &track.Name, &track.Datetime, &track.Author)
		if err != nil {
		    log.Println(err)
		}

		tracks = append(tracks, track)	
	}

	return tracks
	
}

func GettracksIds() []TrackId{
	
	var tracks []TrackId
	rows, err := db.Query("SELECT id FROM track")
	if err != nil {
	    log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		
		var track TrackId
		err = rows.Scan(&track.Id)
		if err != nil {
		    log.Println(err)
		}

		tracks = append(tracks, track)	
	}

	return tracks
}

