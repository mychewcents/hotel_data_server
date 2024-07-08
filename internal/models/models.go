package models

type Hotel struct {
	ID                string       `json:"id"`
	DestinationID     int          `json:"destination_id"`
	Name              string       `json:"name"`
	Location          LocationObj  `json:"location,omitempty"`
	Description       string       `json:"description"`
	Amenities         AmenitiesObj `json:"amenities"`
	Images            ImagesObj    `json:"images"`
	BookingConditions []string     `json:"booking_conditions,omitempty"`
}

type LocationObj struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
}

type AmenitiesObj struct {
	General []string `json:"general,omitempty"`
	Room    []string `json:"room,omitempty"`
}

type ImagesObj struct {
	Rooms     []SingleImageObj `json:"rooms,omitempty"`
	Site      []SingleImageObj `json:"site,omitempty"`
	Amenities []SingleImageObj `json:"amenities,omitempty"`
}

type SingleImageObj struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}
