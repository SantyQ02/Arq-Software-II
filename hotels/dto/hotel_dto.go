package dto

import "github.com/google/uuid"

type Hotel struct {
	HotelID     uuid.UUID `json:"hotel_id,omitempty"`
	AmadeusID	string 	  `json:"amadeus_id" binding:"required"`
	CityCode	string	  `json:"city_code" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	PricePerDay float64   `json:"price_per_day" binding:"required"`
	Photos      Photos    `json:"photos"`
	Amenities   Amenities `json:"amenities"`
	Active      bool      `json:"active"`
}

type Hotels []Hotel
