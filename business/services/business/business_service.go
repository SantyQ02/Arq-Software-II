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
