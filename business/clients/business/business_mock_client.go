package businessClient

import (
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BusinessMockClient struct {
	mock.Mock
}

func (s *BusinessMockClient) InsertHotelMapping(hotelMapping model.HotelMapping) error {
	ret := s.Called(hotelMapping)
	return ret.Error(0)
}

func (s *BusinessMockClient) GetAmadeusIDByHotelID(hotelID uuid.UUID) string {
	ret := s.Called(hotelID)
	return ret.String(0)
}

func (s *BusinessMockClient) GetAmadeusAvailability(amadeusID string, checkInDate time.Time, checkOutDate time.Time) (bool, e.ApiError) {
	ret := s.Called(amadeusID, checkInDate, checkOutDate)
	    if ret.Get(1) == nil {
        return ret.Bool(0),nil
    }
    return ret.Bool(0), ret.Get(1).(e.ApiError)
}