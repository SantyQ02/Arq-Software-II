package dto

import (
	"github.com/google/uuid"
)

type Amenitie struct {
	AmenitieID uuid.UUID `json:"amenitie_id"`
	Title      string    `json:"title"`
}

type Amenities []Amenitie
