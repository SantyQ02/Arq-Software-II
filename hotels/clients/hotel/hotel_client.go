package hotelClient

import (
	"errors"
	"mvc-go/model"
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	Db *mongo.Database 
	
)

func init() {
	HotelClient = &hotelClient{}
}


func (c *hotelClient) GetHotelById(id string) model.Hotel {
	var hotel model.Hotel

	err := Db.Collection("hotels").FindOne(context.TODO(), bson.D{{"HotelID", id}}).Decode(&hotel)
	if err != nil {
		log.Error("")
		return model.Hotel{}
	}

	return hotel
}


func (c *hotelClient) InsertHotel(hotel model.Hotel) model.Hotel {
	_, err := Db.Collection("hotels").InsertOne(context.TODO(), &hotel)

	if err != nil {
		log.Error("")
		return model.Hotel{}
	}
	log.Debug("Hotel Created: ", hotel.HotelID)
	return hotel
}

func (c *hotelClient) UpdateHotel(hotel model.Hotel) model.Hotel {
	_, err := Db.Collection("hotels").UpdateOne(context.TODO(), bson.D{{"HotelID", hotel.HotelID}},bson.D{{"$set",&hotel}})

	if err != nil {
		log.Error(err.Error())
		return model.Hotel{}
	}
	return hotel
}

func (c *hotelClient) DeleteHotel(id string) error {
	_, err := Db.Collection("hotels").DeleteOne(context.TODO(), bson.D{{"HotelID", id}})
	if err != nil {
		log.Debug(id)
		log.Error(err.Error())
		return errors.New(err.Error())
	}
	return nil
}


