package dto

import (
	"github.com/google/uuid"
)

type HotelMapping struct {
	HotelID   uuid.UUID `json:"hotel_id"`
	AmadeusID string    `json:"amadeus_id"`
}

type HotelMappings []HotelMapping
