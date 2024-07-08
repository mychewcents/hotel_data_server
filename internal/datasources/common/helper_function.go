package common

import "github.com/mychewcents/hotel_data_server/internal/models"

func ShouldShowHotel(hotel *models.Hotel, hotelIDsAsMap map[string]bool, destIDAsMap map[int]bool) bool {
	return (hotelIDsAsMap == nil || hotelIDsAsMap[hotel.ID]) && (destIDAsMap == nil || destIDAsMap[hotel.DestinationID])
}
