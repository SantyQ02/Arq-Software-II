package hotelService

import (
	"mvc-go/dto"
	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type HotelMockService struct {
	mock.Mock
}

func (s *HotelMockService) GetHotelById(id uuid.UUID) (dto.Hotel, e.ApiError) {
	ret := s.Called(id)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.Hotel), nil
	}
	return ret.Get(0).(dto.Hotel), ret.Get(1).(e.ApiError)
}
func (s *HotelMockService) GetHotels() (dto.Hotels, e.ApiError) {
	ret := s.Called()
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.Hotels), nil
	}
	return ret.Get(0).(dto.Hotels), ret.Get(1).(e.ApiError)
}
func (s *HotelMockService) InsertHotel(hoteldto dto.Hotel) (dto.Hotel, e.ApiError) {
	ret := s.Called(hoteldto)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.Hotel), nil
	}
	return ret.Get(0).(dto.Hotel), ret.Get(1).(e.ApiError)
}
func (s *HotelMockService) DeleteHotel(id uuid.UUID) e.ApiError {
	ret := s.Called(id)
	if ret.Get(0) == nil {
		return nil
	}
	return ret.Get(0).(e.ApiError)
}
func (s *HotelMockService) UpdateHotel(hoteldto dto.Hotel) (dto.Hotel, e.ApiError) {
	ret := s.Called(hoteldto)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.Hotel), nil
	}
	return ret.Get(0).(dto.Hotel), ret.Get(1).(e.ApiError)
}
func (s *HotelMockService) UploadPhoto(photodto dto.Photo, id uuid.UUID) (dto.Photo, e.ApiError) {
	ret := s.Called(photodto, id)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.Photo), nil
	}
	return ret.Get(0).(dto.Photo), ret.Get(1).(e.ApiError)
}
func (s *HotelMockService) SendMessage(id uuid.UUID, action string) {
	return
}