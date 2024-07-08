package paperfiles

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mychewcents/hotel_data_server/internal/datasources"
	"github.com/mychewcents/hotel_data_server/internal/datasources/common"
	"github.com/mychewcents/hotel_data_server/internal/models"
)

type paperfilesImpl struct{}

const (
	sourceURL = "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/paperflies"
)

func GetHandler() datasources.DataSource {
	return &paperfilesImpl{}
}

func (pfi *paperfilesImpl) GetHotels(hotelIDsAsMap map[string]bool, destIDsAsMap map[int]bool) ([]*models.Hotel, error) {
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

func (pfi *paperfilesImpl) UpdateHotelDetails(hotelsAsMap map[string]*models.Hotel, paperfilesHotels []*models.Hotel) {
	// Empty function because we take this data sources as our base data
}
