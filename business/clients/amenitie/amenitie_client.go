package amenitieClient

import (
	"errors"
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type amenitieClient struct{}

type amenitieClientInterface interface {
	GetAmenitieById(id string) model.Amenitie
	GetAmenities() model.Amenities
	InsertAmenitie(amenitie model.Amenitie) model.Amenitie
	UpdateAmenitie(amenitie model.Amenitie) model.Amenitie
	DeleteAmenitie(id string) error
	LoadAmenities(id string, Amenities model.Amenities) error
	UnloadAmenities(id string, Amenities model.Amenities) error
}

var (
	AmenitieClient amenitieClientInterface
)

func init() {
	AmenitieClient = &amenitieClient{}
}

var Db *gorm.DB

func (c *amenitieClient) GetAmenitieById(id string) model.Amenitie {
	var amenitie model.Amenitie

	Db.First(&amenitie, "amenitie_id = ?", id)
	log.Debug("amenitie: ", amenitie)

	return amenitie
}

func (c *amenitieClient) GetAmenities() model.Amenities {
	var amenities model.Amenities
	result := Db.Find(&amenities)
	if result.Error != nil {
		log.Error("")
		return model.Amenities{}
	}
	log.Debug("Amenities: ", amenities)

	return amenities
}

func (c *amenitieClient) InsertAmenitie(amenitie model.Amenitie) model.Amenitie {
	result := Db.Create(&amenitie)

	if result.Error != nil {
		log.Error("")
		return model.Amenitie{}
	}
	log.Debug("amenitie Created: ", amenitie.AmenitieID)
	return amenitie
}

func (c *amenitieClient) UpdateAmenitie(amenitie model.Amenitie) model.Amenitie {
	result := Db.Save(&amenitie)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return model.Amenitie{}
	}
	return amenitie
}

func (c *amenitieClient) DeleteAmenitie(id string) error {
	var amenitie model.Amenitie
	result := Db.Delete(&amenitie, "amenitie_id = ?", id)
	if result.Error != nil {
		log.Debug(id)
		log.Error(result.Error.Error())
		return errors.New(result.Error.Error())
	}
	return nil
}

func (c *amenitieClient) LoadAmenities(id string, Amenities model.Amenities) error {
	var hotel model.Hotel
	result := Db.First(&hotel, "hotel_id = ?", id)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return errors.New(result.Error.Error())
	}

	for _, amenity := range Amenities {
		result := Db.Model(&hotel).Association("Amenities").Append(&amenity)
		if result.Error != nil {
			log.Error(result.Error.Error())
			return errors.New(result.Error.Error())
		}
	}

	return nil
}

func (c *amenitieClient) UnloadAmenities(id string, Amenities model.Amenities) error {
	var hotel model.Hotel
	result := Db.First(&hotel, "hotel_id = ?", id)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return errors.New(result.Error.Error())
	}

	for _, amenity := range Amenities {
		result := Db.Model(&hotel).Association("Amenities").Delete(&amenity)
		if result.Error != nil {
			log.Error(result.Error.Error())
			return errors.New(result.Error.Error())
		}
	}

	return nil
}
