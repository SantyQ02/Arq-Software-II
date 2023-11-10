package dto

import "github.com/google/uuid"

type BusinessResponse struct {
	HotelID uuid.UUID `json:"hotel_id"`
	Available bool `json:"available"`
}