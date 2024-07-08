package patagonia

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mychewcents/hotel_data_server/internal/datasources"
	"github.com/mychewcents/hotel_data_server/internal/datasources/common"
	"github.com/mychewcents/hotel_data_server/internal/models"
)

type patagoniaImpl struct{}

const (
	sourceURL = "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"
)

func GetHandler() datasources.DataSource {
	return &patagoniaImpl{}
}

func (pi *patagoniaImpl) GetHotels(hotelIDsAsMap map[string]bool, destIDsAsMap map[int]bool) ([]*models.Hotel, error) {
	resp, err := http.Get(sourceURL)

	// Return nil if error/non-200 status is received.
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("non 2xx status code, statusCode=%v", resp.StatusCode))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// Can add logging here
		return nil, err
	}

	hotels := make([]Hotel, 0)

	err = json.Unmarshal(respBody, &hotels)
	if err != nil {
		return nil, err
	}

	finalHotelList := make([]*models.Hotel, 0)

	for _, hotel := range hotels {
		finalHotel := hotel.ConvertToHotel()
		if common.ShouldShowHotel(finalHotel, hotelIDsAsMap, destIDsAsMap) {
			finalHotelList = append(finalHotelList, finalHotel)
		}
	}

	return finalHotelList, nil
}

func (pi *patagoniaImpl) UpdateHotelDetails(hotelsAsMap map[string]*models.Hotel, patagoniaHotels []*models.Hotel) {
	for _, hotel := range patagoniaHotels {
		if h, exists := hotelsAsMap[hotel.ID]; exists {
			h.Location.Latitude = hotel.Location.Latitude
			h.Location.Longitude = hotel.Location.Longitude

			if hotel.Amenities.Room != nil {
				existingAmenities := h.Amenities.Room
				existingAmenitiesAsMap := map[string]bool{}

				for _, amenity := range existingAmenities {
					existingAmenitiesAsMap[strings.ToLower(amenity)] = true
				}

				for _, amenity := range hotel.Amenities.Room {
					if !existingAmenitiesAsMap[strings.ToLower(amenity)] {
						existingAmenities = append(existingAmenities, amenity)
					}
				}

				h.Amenities.Room = existingAmenities
			}

			h.Images.Amenities = hotel.Images.Amenities
			if len(hotel.Images.Rooms) > 0 {
				roomsImages := h.Images.Rooms
				existingRoomImages := map[string]bool{}
				for _, image := range roomsImages {
					existingRoomImages[image.Link] = true
				}

				for _, image := range hotel.Images.Rooms {
					if !existingRoomImages[image.Link] {
						roomsImages = append(roomsImages, image)
					}
				}

				h.Images.Rooms = roomsImages
			}
		} else {
			hotelsAsMap[hotel.ID] = hotel
		}
	}
}
