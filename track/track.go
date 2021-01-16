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
	Track_data 	string 
	Url 			string 
	Datetime 		string
	Updatetime		string
}

type TrackId struct {
	Id				int
}

var connStr = os.Getenv("DATABASE_URL")
var db, err = sql.Open("postgres", connStr)

func GetTrack(track_id int) Track {

	var track Track
	err := db.QueryRow("SELECT * FROM track WHERE id = $1;", track_id).Scan(&track.Id, &track.Name, &track.Track_data, &track.Url, &track.Datetime, &track.Updatetime)
	if err != nil {
	    log.Println(err)
	}

	return Track {
		Id: track.Id,
		Name: track.Name,
		Track_data: track.Track_data,
		Url: track.Url,
		Datetime: track.Datetime,
		Updatetime: track.Updatetime,
	}
}

func GetTracks() []Track {

	var tracks []Track
	rows, err := db.Query("SELECT id, name, datetime FROM track")
	if err != nil {
	    log.Println(err)
	}

	log.Println(rows)
	
	defer rows.Close()

	for rows.Next() {
		
		var track Track
		err = rows.Scan(&track.Id, &track.Name, &track.Datetime)
		if err != nil {
		    log.Println(err)
		}

		tracks = append(tracks, track)	
	}

	return tracks
	
}

func GetTrackIds() []TrackId{
	
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

