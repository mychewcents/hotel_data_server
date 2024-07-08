package paperfiles

import (
	"github.com/mychewcents/hotel_data_server/internal/models"
	"strings"
)

type Hotel struct {
	ID                string       `json:"hotel_id"`
	DestinationID     int          `json:"destination_id"`
	Name              string       `json:"hotel_name"`
	Location          LocationObj  `json:"location"`
	Description       string       `json:"details"`
	Amenities         AmenitiesObj `json:"amenities"`
	Images            ImageObj     `json:"images"`
	BookingConditions []string     `json:"booking_conditions"`
}

type LocationObj struct {
	Address string `json:"address"`
	Country string `json:"country"`
}

type AmenitiesObj struct {
	General []string `json:"general"`
	Room    []string `json:"room"`
}

type ImageObj struct {
	Rooms []SingleImageObj `json:"rooms"`
	Site  []SingleImageObj `json:"site"`
}

type SingleImageObj struct {
	Link        string `json:"link"`
	Description string `json:"caption"`
}

func (pfh *Hotel) ConvertToHotel() *models.Hotel {

	generalAmenities := make([]string, 0, len(pfh.Amenities.General))
	roomAmenities := make([]string, 0, len(pfh.Amenities.Room))

	for _, amenity := range pfh.Amenities.General {
		generalAmenities = append(generalAmenities, strings.ToLower(strings.Trim(amenity, " ")))
	}

	for _, amenity := range pfh.Amenities.Room {
		roomAmenities = append(roomAmenities, strings.ToLower(strings.Trim(amenity, " ")))
	}

	roomsImages := make([]models.SingleImageObj, 0)
	siteImages := make([]models.SingleImageObj, 0)

	for _, roomImage := range pfh.Images.Rooms {
		roomsImages = append(roomsImages, models.SingleImageObj{
			Link:        roomImage.Link,
			Description: roomImage.Description,
		})
	}

	for _, siteImage := range pfh.Images.Site {
		siteImages = append(siteImages, models.SingleImageObj{
			Link:        siteImage.Link,
			Description: siteImage.Description,
		})
	}

	return &models.Hotel{
		ID:            pfh.ID,
		DestinationID: pfh.DestinationID,
		Name:          pfh.Name,
		Location: models.LocationObj{
			Address: pfh.Location.Address,
			Country: pfh.Location.Country,
		},
		Description: pfh.Description,
		Images: models.ImagesObj{
			Rooms: roomsImages,
			Site:  siteImages,
		},
		Amenities: models.AmenitiesObj{
			General: generalAmenities,
			Room:    roomAmenities,
		},
		BookingConditions: pfh.BookingConditions,
	}
}
