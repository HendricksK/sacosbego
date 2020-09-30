package article

import (
	// "encoding/json"
	// "log"
	// "fmt"
)

type Article struct {
	Id				int
	Name 			string 
	Article_data 	string 
	Url 			string 
	Datetime 		string
}

func GetArticle(id int) Article {
	return Article{
		Id:12,
		Name:"Defectors not heroes 'Cycling'",
		Article_data:"qwerty",
		Url:"https://raw.githubusercontent.com/HendricksK/sacos_images/master/defectors_not_heroes_cycling.pdf",
		Datetime:"2020-09-29 20:21:27"}
}

