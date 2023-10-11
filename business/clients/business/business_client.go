package businessClient

import (
	"os"
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

var AmadeusToken string
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

	// AmadeusToken use

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

func getAmadeusToken() {
	clientID := os.Getenv("AMADEUS_API_KEY")
	clientSecret := os.Getenv("AMADEUS_API_SECRET")

	token_url := fmt.Sprintf("https://test.api.amadeus.com/v1/security/oauth2/token?grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret)

	response, err := http.Get(token_url)
	if err != nil {
		fmt.Println("Error en la solicitud:", err)
		return	
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error en la solicitud:", response.StatusCode)
		return
	}

	// Decodificar la respuesta JSON en la estructura AccessTokenResponse
	var tokenResponse AccessTokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	if err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return
	}

	accessToken := tokenResponse.AccessToken
}

type AccessTokenResponse struct {
	
}

