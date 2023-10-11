package businessClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"mvc-go/model"
	"net/http"
	"os"
	"strings"
	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type businessClient struct{}

type businessClientInterface interface {
	InsertHotelMapping(hotelMapping model.HotelMapping) error
	GetAmadeusIDByHotelID(hotelID uuid.UUID) string
	GetAmadeusAvailability(amadeusID string, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError)
}

var (
	BusinessClient businessClientInterface
)

func init() {
	BusinessClient = &businessClient{}
	go getAmadeusToken()
}

var AmadeusToken string
var Db *gorm.DB

func (s *businessClient) InsertHotelMapping(hotelMapping model.HotelMapping) error {
	result := Db.Create(&hotelMapping)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *businessClient) GetAmadeusIDByHotelID(hotelID uuid.UUID) string {
	var hotelMapping model.HotelMapping
	Db.First(&hotelMapping, "hotel_id = ?", hotelID)

	return hotelMapping.AmadeusID
}

func (s *businessClient) GetAmadeusAvailability(amadeusID string, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError) {
	url := fmt.Sprintf("https://test.api.amadeus.com/v3/shopping/hotel-offers?hotelIds=%s&checkInDate=%s&checkOutDate=%s", amadeusID, checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))
	tokenHeader := fmt.Sprintf("Bearer %s", AmadeusToken)

	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, e.NewInternalServerApiError("Error creating request to get Amadeus availability!", errors.New(""))
	}

	req.Header.Set("Authorization", tokenHeader)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error en la solicitud:", err)
		return false, e.NewInternalServerApiError("Error getting response from Amadeus availability!", errors.New(""))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error en la solicitud:", resp.StatusCode)
		return false, e.NewInternalServerApiError("Error getting response from Amadeus availability! (Status-Code)", errors.New(""))
	}

	var availabilityResponse availabilityResponse

	err = json.NewDecoder(resp.Body).Decode(&availabilityResponse)
	if err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return false, nil
	}

	if len(availabilityResponse.Data) == 0 {
		return false, nil
	}

	return availabilityResponse.Data[0].Available, nil
}

func getAmadeusToken() {
	for {
		clientID := os.Getenv("AMADEUS_API_KEY")
		clientSecret := os.Getenv("AMADEUS_API_SECRET")

		tokenURL := "https://test.api.amadeus.com/v1/security/oauth2/token"
		data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret)

		req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := http.Client{}
		response, err := client.Do(req)
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
		var tokenResponse accessTokenResponse
		err = json.NewDecoder(response.Body).Decode(&tokenResponse)
		if err != nil {
			fmt.Println("Error al decodificar la respuesta JSON:", err)
			return
		}

		AmadeusToken = tokenResponse.AccessToken

		log.Info(AmadeusToken)

		sleepTime := 1700 * time.Second
		time.Sleep(sleepTime)
	}
}

type accessTokenResponse struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientID        string `json:"client_id"`
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

type availabilityResponse struct {
	Data []struct {
		Available bool `json:"available" binding:"required"`
	}
}
