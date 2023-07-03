package models 

type Image struct {
	Id 			*int 		`json:"id"`
	EntityId	*int 		`json:"entity_id"`
	Url 		*string 		`json:"url"`
}

type ImageAggregate struct {
	ImageId			*int 		`json:"image_id"`
	Tags 			*string 	`json:"tags"` 
	CreatedAt 		*time.Time  `json:"created_at"`
	UpdatedAt 		*time.Time  `json:"updated_at"`
}

var model = ""
var model_aggregate = ""
var base_url = "https://raw.githubusercontent.com/HendricksK/sacos_images/"

func GetImages(entity_id int) {

}