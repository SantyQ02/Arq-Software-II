package photoClient

import (
	"errors"
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type photoClient struct{}

type photoClientInterface interface {
	GetPhotoById(id string) model.Photo
	GetPhotosByHotelId(hotelId string) model.Photos
	InsertPhoto(photo model.Photo) model.Photo
	UpdatePhoto(photo model.Photo) model.Photo
	DeletePhoto(id string) error
}

var (
	PhotoClient photoClientInterface
)

func init() {
	PhotoClient = &photoClient{}
}

var Db *gorm.DB

func (c *photoClient) GetPhotoById(id string) model.Photo {
	var photo model.Photo

	Db.First(&photo, "photo_id = ?", id)
	log.Debug("photo: ", photo)

	return photo
}

func (c *photoClient) GetPhotosByHotelId(hotelId string) model.Photos {
	var photos model.Photos
	result := Db.Find(&photos, "hotel_id = ?", hotelId)
	if result.Error != nil {
		log.Error("")
		return model.Photos{}
	}
	log.Debug("photos: ", photos)

	return photos
}

func (c *photoClient) InsertPhoto(photo model.Photo) model.Photo {
	result := Db.Create(&photo)

	if result.Error != nil {
		log.Error("")
		return model.Photo{}
	}
	log.Debug("photo Created: ", photo.PhotoID)
	return photo
}

func (c *photoClient) UpdatePhoto(photo model.Photo) model.Photo {
	result := Db.Save(&photo)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return model.Photo{}
	}
	return photo
}

func (c *photoClient) DeletePhoto(id string) error {
	var photo model.Photo
	result := Db.Delete(&photo, "photo_id = ?", id)
	if result.Error != nil {
		log.Debug(id)
		log.Error(result.Error.Error())
		return errors.New(result.Error.Error())
	}
	return nil
}
