package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Amenitie struct {
	AmenitieID uuid.UUID `gorm:"type:char(36);primary_key"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Hotels     Hotels    `gorm:"many2many:hotel_amenities;"`
}

type Amenities []Amenitie

func (amenitie *Amenitie) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("AmenitieID", uuid.New())
}
