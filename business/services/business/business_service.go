package businessService

import (
	"errors"
	businessClient "mvc-go/clients/business"
	hotelService "mvc-go/services/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
)

type businessService struct{}

type businessServiceInterface interface {
	CheckAvailability(id uuid.UUID) (bool, e.ApiError)
	MapHotel(hotelMappingDto dto.HotelMapping) (dto.HotelMapping, e.ApiError)
}

var (
	BusinessService businessServiceInterface
)

func init() {
	BusinessService = &businessService{}
}

func (s *businessService) CheckAvailability(id uuid.UUID) (bool, e.ApiError) {
	return false, nil
}

func (s *businessService) MapHotel(hotelMappingDto dto.HotelMapping) (dto.HotelMapping, e.ApiError) {
	_, err := hotelService.HotelService.GetHotelById(hotelMappingDto.HotelID)
	if err != nil{
		return dto.HotelMapping{}, err
	}
	
	hotelMapping := model.HotelMapping{
		HotelID: hotelMappingDto.HotelID,
		AmadeusID: hotelMappingDto.AmadeusID,
	}

	businessClient.BusinessClient.InsertHotelMapping(hotelMapping)
	
	return hotelMappingDto, nil
}
