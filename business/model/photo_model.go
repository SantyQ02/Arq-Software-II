package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	PhotoID uuid.UUID `gorm:"type:char(36);primary_key"`
	Url     string    `gorm:"type:varchar(255);not null"`
	HotelID uuid.UUID `gorm:"type:char(36);not null"`
}

type Photos []Photo

func (photo *Photo) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("PhotoID", uuid.New())
}
