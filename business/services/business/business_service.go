package businessService

import (
	"fmt"
	"mvc-go/cache"
	businessClient "mvc-go/clients/business"
	userClient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	// "os"
)

type businessService struct{}

type businessServiceInterface interface {
	CheckAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError)
	MapHotel(hotelMappingDto dto.HotelMapping) (dto.HotelMapping, e.ApiError)
	CheckAdmin(userID uuid.UUID) (bool, e.ApiError)
}

var (
	BusinessService businessServiceInterface
)

func init() {
	BusinessService = &businessService{}
}

func (s *businessService) CheckAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError) {
	cacheKey := fmt.Sprintf("%s/%s/%s", id.String(), checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))

	availability := cache.Get(cacheKey)
	if availability != nil {
		return bytesToBool(availability), nil
	}

	amadeusID := businessClient.BusinessClient.GetAmadeusIDByHotelID(id)
	if amadeusID == "" {
		return false, e.NewNotFoundApiError("Hotel not found!")
	}

	available, err := businessClient.BusinessClient.GetAmadeusAvailability(amadeusID, checkInDate, checkOutDate)
	if err != nil {
		return false, err
	}

	cache.SetWithExpiration(cacheKey, boolToBytes(available), 10) // Cache Expiration 10 seconds

	return available, nil
}

func (s *businessService) MapHotel(hotelMappingDto dto.HotelMapping) (dto.HotelMapping, e.ApiError) {
	hotelMapping := model.HotelMapping{
		HotelID:   hotelMappingDto.HotelID,
		AmadeusID: hotelMappingDto.AmadeusID,
	}

	err := businessClient.BusinessClient.InsertHotelMapping(hotelMapping)
	if err != nil {
		return dto.HotelMapping{}, e.NewBadRequestApiError("AmadeusID already mapped to an existing HotelID or HotelID already mapped to an AmadeusID.")
	}

	return hotelMappingDto, nil
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

func (s *businessService) CheckAdmin(userID uuid.UUID) (bool, e.ApiError) {
	user := userClient.UserClient.GetUserById(userID.String())
	if user.UserID == uuid.Nil {
		return false, e.NewNotFoundApiError("User not found")
	}

	if user.Role != "admin" {
		return false, nil
	}

	return true, nil
}
