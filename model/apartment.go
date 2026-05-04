package model

type Apartment struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	DescriptionShort string   `json:"shortDescription"`
	Description      string   `json:"description"`
	Price            int      `json:"price"`
	Location         string   `json:"location"`
	Type             string   `json:"type"`
	Size             int      `json:"size"`
	Beds             int      `json:"beds"`
	Baths            int      `json:"baths"`
	Amenities        []string `json:"amenities"`
	Images           []string `json:"images"`
	FloorPlan        string   `json:"floorPlan"`
	Maps             string   `json:"maps"`
}

type Apartments struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Price  int      `json:"price"`
	Images []string `json:"images"`
	Size   int      `json:"size"`
	Beds   int      `json:"beds"`
	Baths  int      `json:"baths"`
}
