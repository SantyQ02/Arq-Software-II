package hotelClient

import (
	"errors"
	"mvc-go/model"
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"bytes"
	"encoding/json"
	"net/http"
	"github.com/google/uuid"

)

type hotelClient struct{}

type hotelClientInterface interface {
	GetHotelById(id string) model.Hotel
	InsertHotel(hotel model.Hotel) model.Hotel
	UpdateHotel(hotel model.Hotel) model.Hotel
	DeleteHotel(id string) error
	HotelMapping(hotel model.Hotel) error
}

var (
	HotelClient hotelClientInterface
	Db *mongo.Database 
	
)

func init() {
	HotelClient = &hotelClient{}
}

type Bodystruct struct {
    HotelID uuid.UUID `json:"hotel_id,omitempty"`
	AmadeusID  string `json:"amadeus_id,omitempty"`
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


func (c *hotelClient) HotelMapping(hotel model.Hotel) error {

	businessURL := "http://business:8080/api/business/mapping-hotel"

	hotelID,_ := uuid.Parse(hotel.HotelID)

	// JSON body
	body := Bodystruct{
		HotelID:   hotelID,
		AmadeusID: hotel.AmadeusID,
	}

	jsonData, _ := json.Marshal(body)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", businessURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	_, er := client.Do(r)

	if er != nil {
		return er
	}

	return nil
}
