package businessClient

import (
	"mvc-go/model"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type businessClient struct{}

type businessClientInterface interface {
	InsertHotelMapping(hotelMapping model.HotelMapping)
	GetAmadeusIDByHotelID(hotelID uuid.UUID) string
}

var (
	BusinessClient businessClientInterface
)

func init() {
	BusinessClient = &businessClient{}
}

var Db *gorm.DB

func (s *businessClient) InsertHotelMapping(hotelMapping model.HotelMapping) {
	result := Db.Create(&hotelMapping)
	if result.Error != nil {
		log.Error("")
	}
}

func (s *businessClient) GetAmadeusIDByHotelID(hotelID uuid.UUID) string {
	var hotelMapping model.HotelMapping
	Db.First(&hotelMapping, "hotel_id = ?", hotelID)

	return hotelMapping.AmadeusID
}
