package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID           uuid.UUID `gorm:"type:char(36);primary_key"`
	FirstName        string    `gorm:"type:varchar(255);not null"`
	LastName         string    `gorm:"type:varchar(255);not null"`
	Email            string    `gorm:"unique;not null"`
	UserName         string    `gorm:"unique;not null"`
	Password         string    `gorm:"not null"`
	Role             string    `gorm:"type:varchar(255);not null"`
	VerificationCode string
	Verified         bool      `gorm:"not null"`
	Active           bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
	Bookings         Bookings  `gorm:"foreignKey:UserID"`
}

type Users []User

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("UserID", uuid.New())
}
