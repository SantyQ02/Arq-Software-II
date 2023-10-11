package client

import (
	"mvc-go/dto"
	"net/http"
	"github.com/google/uuid"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	// log "github.com/sirupsen/logrus"
)

type businessClient struct{}

type businessClientInterface interface {
	GetHotelAvailability(id uuid.UUID, checkInDate string, checkOutDate string) (dto.BusinessResponse, error)
}

var (
	BusinessClient businessClientInterface
)

func init() {
	BusinessClient = &businessClient{}
}

func (c *businessClient) GetHotelAvailability(id uuid.UUID, checkInDate string, checkOutDate string) (dto.BusinessResponse, error) {

	BUSINESS_SERVICE_URL := os.Getenv("BUSINESS_SERVICE_URL") 
	hotelId := id.String()
	url := fmt.Sprintf("%s/api/availability/%s?checkInDate=%s&checkOutDate=%s", BUSINESS_SERVICE_URL, hotelId, checkInDate, checkOutDate)
	
	response, err := http.Get(url)
	if err != nil {
		return dto.BusinessResponse{}, err
	}

	// Validate API Error
	if response.StatusCode != http.StatusOK {
		return dto.BusinessResponse{}, err
	}

	// Read response payload bytes
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dto.BusinessResponse{}, err
	}

	// Convert bytes to custom struct
	var businessRes dto.BusinessResponse
	err = json.Unmarshal(bytes, &businessRes)
	if err != nil {
		return dto.BusinessResponse{}, err
	}
	return businessRes, nil
}

