package entity

type Chars struct {
	Status      string   `json:"status,omitempty"`
	Name        string   `json:"name,omitempty"`
	Image       string   `json:"image,omitempty"`
	Image_manga string   `json:"image_manga,omitempty"`
	SplitedName []string `json:"splitedName,omitempty"`
	Bounty      string   `json:"bounty,omitempty"`
	_id         string   `json:"_id,omitempty"`
}

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

type SpecificChar struct {
	Name string `json:"name,omitempty"`
}
