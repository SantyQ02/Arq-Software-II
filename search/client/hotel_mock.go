package client

import (
	"mvc-go/dto"

	"github.com/stretchr/testify/mock"
	"github.com/google/uuid"
)

type HotelMockClient struct {
	mock.Mock
}

func (c *HotelMockClient) GetHotel(id uuid.UUID) (dto.Hotel, error) {
	ret := c.Called(id)
	return ret.Get(0).(dto.Hotel), ret.Error(1)
}