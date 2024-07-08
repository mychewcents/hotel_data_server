package datasources

import (
	"github.com/mychewcents/hotel_data_server/internal/models"
)

type DataSource interface {
	GetHotels(hotelIDsAsMap map[string]bool, destIDsAsMap map[int]bool) ([]*models.Hotel, error)
	UpdateHotelDetails(hotelsAsMap map[string]*models.Hotel, hotels []*models.Hotel)
}
