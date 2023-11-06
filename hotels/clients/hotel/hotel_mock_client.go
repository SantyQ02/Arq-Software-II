package hotelClient

import (
	"mvc-go/model"

	"github.com/stretchr/testify/mock"
)

type HotelMockClient struct {
	mock.Mock
}

func (c *HotelMockClient) GetHotelById(id string) model.Hotel {
	ret := c.Called(id)
	return ret.Get(0).(model.Hotel)
}

func (c *HotelMockClient) InsertHotel(hotel model.Hotel) model.Hotel {
	ret := c.Called(hotel)
	return ret.Get(0).(model.Hotel)
}

func (c *HotelMockClient) UpdateHotel(hotel model.Hotel) model.Hotel {
	ret := c.Called(hotel)
	return ret.Get(0).(model.Hotel)
}

func (c *HotelMockClient) DeleteHotel(id string) error {
	ret := c.Called(id)
	return ret.Error(0)
}