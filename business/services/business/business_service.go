package businessService

import (
	"encoding/json"
	"errors"
	"fmt"
	"mvc-go/cache"
	businessClient "mvc-go/clients/business"
	"mvc-go/dto"
	"mvc-go/model"
	hotelService "mvc-go/services/hotel"
	"net/http"
	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	// "os"
	// "github.com/joho/godotenv"
)

type businessService struct{}

type businessServiceInterface interface {
	CheckAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError)
	MapHotel(hotelMappingDto dto.HotelMapping) (dto.HotelMapping, e.ApiError)
	HotelIDToAmadeusID(hotelID uuid.UUID) (string, e.ApiError)
}

var (
	BusinessService businessServiceInterface
)

func init() {
	BusinessService = &businessService{}
}

type availabilityResponse struct {
	Available bool `json:"available" binding:"required"`
}

func (s *businessService) CheckAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError) {
	cacheKey := fmt.Sprintf("%s/%s/%s", id.String(), checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))

	log.Info(cacheKey) // Debug

	availability := cache.Get(cacheKey)
	if availability != nil {
		return bytesToBool(availability), nil
	}

	amadeusID, er := BusinessService.HotelIDToAmadeusID(id)
	if er != nil {
		return false, er
	}

	url := fmt.Sprintf("https://test.api.amadeus.com/v3/shopping/hotel-offers?hotelIds=%s&checkInDate=%s&checkOutDate=%s", amadeusID, checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))

	// Add Amadeus TOKEN Generator
	// Cargar variables de entorno desde el archivo .env.docker
	// err := godotenv.Load(".env.docker")
	// if err != nil {
	// 	fmt.Println("Error al cargar el archivo .env.docker")
	// 	return
	// }

	// // Obtener client_id y client_secret del archivo .env.docker
	// clientID := os.Getenv("AMADEUS_API_KEY")
	// clientSecret := os.Getenv("AMADEUS_API_SECRET")

	// token_url := fmt.Sprintf("https://test.api.amadeus.com/v1/security/oauth2/token?grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret)

	// response, err := http.Get(token_url)
	// if err != nil {
	// 	fmt.Println("Error en la solicitud:", err)
	// 	return
	// }
	// defer response.Body.Close()

	// if response.StatusCode != http.StatusOK {
	// 	fmt.Println("Error en la solicitud:", response.StatusCode)
	// 	return
	// }

	// // Decodificar la respuesta JSON en la estructura AccessTokenResponse
	// var tokenResponse AccessTokenResponse
	// err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	// if err != nil {
	// 	fmt.Println("Error al decodificar la respuesta JSON:", err)
	// 	return
	// }

	// accessToken := tokenResponse.AccessToken

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al hacer la solicitud:", err)
		return false, e.NewInternalServerApiError("Amadeus se cae a los pedazos y no devuelve nada", errors.New(""))
	}

	defer resp.Body.Close()

	var availabilityResponse availabilityResponse

	err = json.NewDecoder(resp.Body).Decode(&availabilityResponse)
	if err != nil {
		availabilityResponse.Available = false
	}

	cache.SetWithExpiration(cacheKey, boolToBytes(availabilityResponse.Available), 10) // Cache Expiration 10 seconds

	return availabilityResponse.Available, nil
}

func (s *businessService) MapHotel(hotelMappingDto dto.HotelMapping) (dto.HotelMapping, e.ApiError) {
	_, err := hotelService.HotelService.GetHotelById(hotelMappingDto.HotelID)
	if err != nil {
		return dto.HotelMapping{}, err
	}

	hotelMapping := model.HotelMapping{
		HotelID:   hotelMappingDto.HotelID,
		AmadeusID: hotelMappingDto.AmadeusID,
	}

	businessClient.BusinessClient.InsertHotelMapping(hotelMapping)

	return hotelMappingDto, nil
}

func (s *businessService) HotelIDToAmadeusID(hotelID uuid.UUID) (string, e.ApiError) {
	amadeusID := businessClient.BusinessClient.GetAmadeusIDByHotelID(hotelID)
	if amadeusID == "" {
		return "", e.NewNotFoundApiError("Hotel not found!")
	}

	return amadeusID, nil
}

func bytesToBool(data []byte) bool {
	if len(data) > 0 && data[0] != 0 {
		return true
	}
	return false
}

func boolToBytes(input bool) []byte {
	if input {
		return []byte{1}
	}
	return []byte{0}
}
