package patagonia

import (
	"github.com/mychewcents/hotel_data_server/internal/models"
	"strings"
)

type Hotel struct {
	ID            string   `json:"id"`
	DestinationID int      `json:"destination"`
	Name          string   `json:"name"`
	Latitude      float64  `json:"lat"`
	Longitude     float64  `json:"lng"`
	Address       string   `json:"address"`
	Description   string   `json:"info"`
	Amenities     []string `json:"amenities"`
	Images        ImageObj `json:"images"`
}

type ImageObj struct {
	Rooms     []SingleImageObj `json:"rooms"`
	Amenities []SingleImageObj `json:"amenities"`
}

type SingleImageObj struct {
	Link        string `json:"url"`
	Description string `json:"description"`
}

func (ph *Hotel) ConvertToHotel() *models.Hotel {

	roomAmenities := make([]string, 0, len(ph.Amenities))
	for _, amenity := range ph.Amenities {
		roomAmenities = append(roomAmenities, strings.ToLower(strings.Trim(amenity, " ")))
	}

	roomsImages := make([]models.SingleImageObj, 0, len(ph.Images.Rooms))
	amenitiesImages := make([]models.SingleImageObj, 0, len(ph.Images.Amenities))

	for _, room := range ph.Images.Rooms {
		roomsImages = append(roomsImages, models.SingleImageObj{
			Link:        room.Link,
			Description: room.Description,
		})
	}

	for _, amenity := range ph.Images.Amenities {
		amenitiesImages = append(amenitiesImages, models.SingleImageObj{
			Link:        amenity.Link,
			Description: amenity.Description,
		})
	}

	return &models.Hotel{
		ID:            ph.ID,
		DestinationID: ph.DestinationID,
		Name:          ph.Name,
		Location: models.LocationObj{
			Latitude:  ph.Latitude,
			Longitude: ph.Longitude,
			Address:   ph.Address,
		},
		Description: ph.Description,
		Images: models.ImagesObj{
			Rooms:     roomsImages,
			Amenities: amenitiesImages,
		},
		Amenities: models.AmenitiesObj{
			Room: roomAmenities,
		},
	}
}
