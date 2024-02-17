package entity

type CharsMongo struct {
	Status      string   `bson:"status,omitempty"`
	Name        string   `bson:"name,omitempty"`
	Image       string   `bson:"image,omitempty"`
	Image_manga string   `bson:"image_manga,omitempty"`
	SplitedName []string `bson:"splitedName,omitempty"`
	Bounty      string   `bson:"bounty,omitempty"`
	Appearance  string   `bson:"appearance,omitempty"`
	_Id         string   `bson:"_id,omitempty"`
}
