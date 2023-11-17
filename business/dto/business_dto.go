package dto

import (
	"github.com/google/uuid"
)

type HotelMapping struct {
	HotelID   uuid.UUID `json:"hotel_id" binding:"required"`
	AmadeusID string    `json:"amadeus_id" binding:"required"`
}

type HotelMappings []HotelMapping
