package client

import (
	"mvc-go/dto"
	"net/http"
	"github.com/google/uuid"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"time"
	"errors"
	// log "github.com/sirupsen/logrus"
)

type businessClient struct{}

type businessClientInterface interface {
	GetHotelAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (dto.BusinessResponse, error)
}

var (
	BusinessClient businessClientInterface
)

func init() {
	BusinessClient = &businessClient{}
}

func (c *businessClient) GetHotelAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (dto.BusinessResponse, error) {

	checkInDateStr := checkInDate.Format("2006-01-02")
	checkOutDateStr := checkOutDate.Format("2006-01-02")

	BUSINESS_SERVICE_URL := os.Getenv("BUSINESS_SERVICE_URL") 
	hotelId := id.String()
	url := fmt.Sprintf("%s/api/business/availability/%s?checkInDate=%s&checkOutDate=%s", BUSINESS_SERVICE_URL, hotelId, checkInDateStr, checkOutDateStr)
	
	response, err := http.Get(url)
	if err != nil {
		return dto.BusinessResponse{}, err
	}

	// Validate API Error
	if response.StatusCode != http.StatusOK {
		return dto.BusinessResponse{}, errors.New("Error checking availability")
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

