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

type SpecificChar struct {
	Name string `json:"name,omitempty"`
}
