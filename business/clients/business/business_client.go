package businessClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"mvc-go/model"
	"net/http"
	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type businessClient struct{}

type businessClientInterface interface {
	InsertHotelMapping(hotelMapping model.HotelMapping)
	GetAmadeusIDByHotelID(hotelID uuid.UUID) string
	GetAmadeusAvailability(amadeusID string, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError)
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

type availabilityResponse struct {
	Available bool `json:"available" binding:"required"`
}

func (s *businessClient) GetAmadeusAvailability(amadeusID string, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError) {
	url := fmt.Sprintf("https://test.api.amadeus.com/v3/shopping/hotel-offers?hotelIds=%s&checkInDate=%s&checkOutDate=%s", amadeusID, checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud:", err)
		return false, e.NewInternalServerApiError("Amadeus se cae a los pedazos y no devuelve nada", errors.New(""))
	}

	defer resp.Body.Close()

	var availabilityResponse availabilityResponse

	err = json.NewDecoder(resp.Body).Decode(&availabilityResponse)
	if err != nil {
		return false, nil
	}

	return availabilityResponse.Available, nil
}
