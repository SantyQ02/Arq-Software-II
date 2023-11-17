package client

import (
	"mvc-go/dto"

	"github.com/stretchr/testify/mock"
	"github.com/google/uuid"
	"time"
)

type BusinessMockClient struct {
	mock.Mock
}

func (c *BusinessMockClient) GetHotelAvailability(id uuid.UUID, checkInDate time.Time, checkOutDate time.Time) (dto.BusinessResponse, error) {
	ret := c.Called(id, checkInDate, checkOutDate)
	return ret.Get(0).(dto.BusinessResponse), ret.Error(1)
}