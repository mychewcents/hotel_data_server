package models

type GetHotelsRequest struct {
	HotelIDs       []string `json:"hotel_ids" url:"hotel_ids"`
	DestinationIDs []int    `json:"destination_ids" url:"destination_ids"`
}
