package dto

import "github.com/google/uuid"

type Amenity struct {
	AmenityID uuid.UUID 	 `json:"amenity_id"`
	Title      string    `json:"title"`
}

type Amenities []Amenity