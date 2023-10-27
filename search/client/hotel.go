package client

import (
	// "errors"
	"mvc-go/dto"
	"net/http"
	"github.com/google/uuid"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

	// log "github.com/sirupsen/logrus"
)

type hotelClient struct{}

type hotelClientInterface interface {
	GetHotel(id uuid.UUID) (dto.Hotel, error)
}

var (
	HotelClient hotelClientInterface
)

func init() {
	HotelClient = &hotelClient{}
}

func (c *hotelClient) GetHotel(id uuid.UUID) (dto.Hotel, error) {

	HOTEL_SERVICE_URL := os.Getenv("HOTEL_SERVICE_URL") 
	hotelId := id.String()
	url := fmt.Sprintf("%s/api/hotel/%s", HOTEL_SERVICE_URL, hotelId)
	
	response, err := http.Get(url)
	if err != nil {
		return dto.Hotel{}, err
	}

	// Validate API Error
	if response.StatusCode != http.StatusOK {
		return dto.Hotel{}, err
	}

	// Read response payload bytes
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.Hotel{}, err
	}

	// Convert bytes to custom struct
	var hotelRes dto.HotelResponse
	err = json.Unmarshal(bytes, &hotelRes)
	if err != nil {
		return dto.Hotel{}, err
	}
	return hotelRes.Hotel, nil
}