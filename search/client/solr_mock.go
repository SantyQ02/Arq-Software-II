package client

import (
	"mvc-go/dto"

	"github.com/stretchr/testify/mock"
)

type SolrMockClient struct {
	mock.Mock
}

func (c *SolrMockClient) SearchHotels(city string) ([]dto.Hotel, error) {
	ret := c.Called(city)
	return ret.Get(0).([]dto.Hotel), ret.Error(1)
}

func (c *SolrMockClient) AddOrUpdateHotel(hotelDto dto.Hotel) error {
	ret := c.Called(hotelDto)
	return ret.Error(0)
}

func (c *SolrMockClient) EmptyCollection() error {
	ret := c.Called()
	return ret.Error(0)
}