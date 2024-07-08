package acme

import (
	"strings"

	"github.com/mychewcents/hotel_data_server/internal/models"
)

type Hotel struct {
	ID            string `json:"Id"`
	DestinationID int    `json:"DestinationId"`
	Name          string `json:"Name"`
	//Latitude      float64  `json:"Latitude"`
	//Longitude     float64  `json:"Longitude"`
	Address     string   `json:"Address"`
	City        string   `json:"City"`
	Country     string   `json:"Country"`
	PostalCode  string   `json:"PostalCode"`
	Description string   `json:"Description"`
	Facilities  []string `json:"Facilities"`
}

func (ah *Hotel) ConvertToHotel() *models.Hotel {

	generalAmenities := make([]string, 0, len(ah.Facilities))
	for _, facility := range ah.Facilities {
		generalAmenities = append(generalAmenities, strings.ToLower(strings.Trim(facility, " ")))
	}

	return &models.Hotel{
		ID:            ah.ID,
		DestinationID: ah.DestinationID,
		Name:          ah.Name,
		Location: models.LocationObj{
			//Latitude:  ah.Latitude,
			//Longitude: ah.Longitude,
			Address: ah.Address,
			City:    ah.City,
			Country: ah.Country,
		},
		Description: strings.Trim(ah.Description, ""),
		Amenities: models.AmenitiesObj{
			General: generalAmenities,
		},
	}
}
