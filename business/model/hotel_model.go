package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Hotel struct {
	HotelID     uuid.UUID `gorm:"type:char(36);primary_key"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Rooms       uint      `gorm:"not null"`
	PricePerDay float64   `gorm:"not null"`
	Bookings    Bookings  `gorm:"foreignKey:HotelID"`
	Photos      Photos    `gorm:"foreignKey:HotelID"`
	Amenities   Amenities `gorm:"many2many:hotel_amenities;"`
	Active      bool      `gorm:"not null"`
}

type Hotels []Hotel

func (hotel *Hotel) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("HotelID", uuid.New())
}
