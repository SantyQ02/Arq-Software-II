package hotelClient

import (
	"errors"
	"mvc-go/model"
	"mvc-go/db"
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type hotelClient struct{}

type hotelClientInterface interface {
	GetHotelById(id string) model.Hotel
	InsertHotel(hotel model.Hotel) model.Hotel
	UpdateHotel(hotel model.Hotel) model.Hotel
	DeleteHotel(id string) error
}

var (
	HotelClient hotelClientInterface
)

func init() {
	HotelClient = &hotelClient{}
}


func (c *hotelClient) GetHotelById(id string) model.Hotel {
	var hotel model.Hotel
	db:=db.MongoDb

	err := db.Collection("hotels").FindOne(context.TODO(), bson.D{{"HotelID", id}}).Decode(&hotel)
	if err != nil {
		log.Error("")
		return model.Hotel{}
	}

	return hotel
}


func (c *hotelClient) InsertHotel(hotel model.Hotel) model.Hotel {
	db:=db.MongoDb
	_, err := db.Collection("hotels").InsertOne(context.TODO(), &hotel)

	if err != nil {
		log.Error("")
		return model.Hotel{}
	}
	log.Debug("Hotel Created: ", hotel.HotelID)
	return hotel
}

func (c *hotelClient) UpdateHotel(hotel model.Hotel) model.Hotel {
	db:=db.MongoDb
	_, err := db.Collection("hotels").UpdateOne(context.TODO(), bson.D{{"HotelID", hotel.HotelID}},bson.D{{"$set",&hotel}})

	if err != nil {
		log.Error(err.Error())
		return model.Hotel{}
	}
	return hotel
}

func (c *hotelClient) DeleteHotel(id string) error {
	db:=db.MongoDb
	_, err := db.Collection("hotels").DeleteOne(context.TODO(), bson.D{{"HotelID", id}})
	if err != nil {
		log.Debug(id)
		log.Error(err.Error())
		return errors.New(err.Error())
	}
	return nil
}


