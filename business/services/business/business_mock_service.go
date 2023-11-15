package businessService

import (
	"mvc-go/dto"
	e "mvc-go/utils/errors"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BusinessMockService struct {
	mock.Mock
}

func (s *BusinessMockService) CheckAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError) {
	ret := s.Called(id, checkInDate, checkOutDate)
	if ret.Get(1) == nil {
		return ret.Bool(0), nil
	}
	return ret.Bool(0), ret.Get(1).(e.ApiError)
}

func (s *BusinessMockService) MapHotel(hotelMappingDto dto.HotelMapping) (dto.HotelMapping, e.ApiError) {
	ret := s.Called(hotelMappingDto)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.HotelMapping), nil
	}
	return ret.Get(0).(dto.HotelMapping), ret.Get(1).(e.ApiError)
}

func (s *BusinessMockService) HotelIDToAmadeusID(hotelID uuid.UUID) (string, e.ApiError) {
	ret := s.Called(hotelID)
	if ret.Get(1) == nil {
		return ret.String(0), nil
	}
	return ret.String(0), ret.Get(1).(e.ApiError)
}

func (s *BusinessMockService) CheckAdmin(userID uuid.UUID) (bool, e.ApiError) {
	ret := s.Called(userID)
	if ret.Get(1) == nil {
		return ret.Bool(0), nil
	}
	return ret.Bool(0), ret.Get(1).(e.ApiError)
}
