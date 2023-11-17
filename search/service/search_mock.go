package service

import (
	"mvc-go/dto"
	e "mvc-go/utils/errors"
	"github.com/stretchr/testify/mock"
	"time"
)

type SearchMockService struct {
	mock.Mock
}

func (c *SearchMockService) Search(city string, checkInDate time.Time, checkOutDate time.Time) ([]dto.Hotel, e.ApiError){
	ret := c.Called(city, checkInDate, checkOutDate)
	if ret.Get(1) == nil {
		return ret.Get(0).([]dto.Hotel),nil
	}
	return ret.Get(0).([]dto.Hotel),ret.Get(1).(e.ApiError)
}