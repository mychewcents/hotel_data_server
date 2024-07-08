package controller

import (
	"github.com/mychewcents/hotel_data_server/internal/datasources"
	"github.com/mychewcents/hotel_data_server/internal/datasources/acme"
	"github.com/mychewcents/hotel_data_server/internal/datasources/paperfiles"
	"github.com/mychewcents/hotel_data_server/internal/datasources/patagonia"
	"github.com/mychewcents/hotel_data_server/internal/models"
)

type controllerImpl struct {
	acmeSource       datasources.DataSource
	patagoniaSource  datasources.DataSource
	paperfilesSource datasources.DataSource
}

var controllerObj *controllerImpl

func init() {
	controllerObj = &controllerImpl{}

	acmeHotels := acme.GetHandler()
	if acmeHotels != nil {
		controllerObj.acmeSource = acmeHotels
	}

	patagoniaHotels := patagonia.GetHandler()
	if patagoniaHotels != nil {
		controllerObj.patagoniaSource = patagoniaHotels
	}

	paperfilesHotels := paperfiles.GetHandler()
	if paperfilesHotels != nil {
		controllerObj.paperfilesSource = paperfilesHotels
	}
}

func GetHotels(req *models.GetHotelsRequest) ([]*models.Hotel, error) {

	hotelsAsMap := map[string]*models.Hotel{}
	var filterHotelIDsAsMap map[string]bool
	var filterDestinationIDsAsMap map[int]bool

	if req.HotelIDs != nil && len(req.HotelIDs) > 0 {
		filterHotelIDsAsMap = make(map[string]bool)
		for _, hotelID := range req.HotelIDs {
			filterHotelIDsAsMap[hotelID] = true
		}
	}

	if req.DestinationIDs != nil && len(req.DestinationIDs) > 0 {
		filterDestinationIDsAsMap = make(map[int]bool)
		for _, destinationID := range req.DestinationIDs {
			filterDestinationIDsAsMap[destinationID] = true
		}
	}

	// Taking paperfiles response as a base for our hotels because of the completeness of data across hotels as observed from the dataset
	if controllerObj.paperfilesSource != nil {
		paperfilesHotels, err := controllerObj.paperfilesSource.GetHotels(filterHotelIDsAsMap, filterDestinationIDsAsMap)
		if err != nil {
			return nil, err
		}
		if paperfilesHotels != nil && len(paperfilesHotels) > 0 {
			for _, hotel := range paperfilesHotels {
				hotelsAsMap[hotel.ID] = hotel
			}
		}
	}

	if controllerObj.patagoniaSource != nil {
		patagoniaHotels, err := controllerObj.patagoniaSource.GetHotels(filterHotelIDsAsMap, filterDestinationIDsAsMap)
		if err != nil {
			return nil, err
		}
		if patagoniaHotels != nil && len(patagoniaHotels) > 0 {
			controllerObj.patagoniaSource.UpdateHotelDetails(hotelsAsMap, patagoniaHotels)
		}
	}

	if controllerObj.acmeSource != nil {
		acmeHotels, err := controllerObj.acmeSource.GetHotels(filterHotelIDsAsMap, filterDestinationIDsAsMap)
		if err != nil {
			return nil, err
		}
		if acmeHotels != nil && len(acmeHotels) > 0 {
			controllerObj.acmeSource.UpdateHotelDetails(hotelsAsMap, acmeHotels)
		}
	}

	hotels := make([]*models.Hotel, 0)

	for _, hotel := range hotelsAsMap {
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
