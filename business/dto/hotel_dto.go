package dto

import (
	"github.com/google/uuid"
)

type Hotel struct {
	HotelID     uuid.UUID `json:"hotel_id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Rooms       uint      `json:"rooms" binding:"required"`
	PricePerDay float64   `json:"price_per_day" binding:"required"`
	Photos      Photos    `json:"photos"`
	Amenities   Amenities `json:"amenities"`
	Active      bool      `json:"active,omitempty"`
}

type Hotels []Hotel
