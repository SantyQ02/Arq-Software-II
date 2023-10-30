package model

import (
	"github.com/google/uuid"
)

type HotelMapping struct {
	HotelID   uuid.UUID `gorm:"type:char(36);primary_key"`
	AmadeusID string    `gorm:"type:char(36);not null;unique"`
}

type HotelMappings []HotelMapping
