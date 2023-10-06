package dto

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	BookingID uuid.UUID `json:"booking_id"`
	Total     float64   `json:"total" binding:"required"`
	DateIn    time.Time `json:"date_in" binding:"required"`
	DateOut   time.Time `json:"date_out" binding:"required"`
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	HotelID   uuid.UUID `json:"hotel_id" binding:"required"`
	Active    bool      `json:"active"`
}

type SetActive struct {
	Active bool `json:"active"`
}

type Bookings []Booking

type CheckAvailability struct {
	HotelID uuid.UUID
	DateIn  time.Time `json:"date_in" binding:"required"`
	DateOut time.Time `json:"date_out" binding:"required"`
}

type CheckAvailabilities []CheckAvailability
