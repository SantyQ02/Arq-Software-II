package dto

import "github.com/google/uuid"

type Photo struct {
	PhotoID uuid.UUID `json:"photo_id"`
	Url     string `json:"url" binding:"required"`
}

type Photos []Photo
