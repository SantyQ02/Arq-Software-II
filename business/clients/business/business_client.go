package businessClient

import (
	"mvc-go/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type businessClient struct{}

type businessClientInterface interface {
	InsertHotelMapping(hotelMapping model.HotelMapping)
}

var (
	BusinessClient businessClientInterface
)

func init() {
	BusinessClient = &businessClient{}
}

var Db *gorm.DB

func (s* businessClient) InsertHotelMapping(hotelMapping model.HotelMapping) {
	result := Db.Create(&hotelMapping)
	if result.Error != nil {
		log.Error("")
	}
}