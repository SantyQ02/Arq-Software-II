package client

import (
	"mvc-go/dto"

	"github.com/stretchr/testify/mock"
)

type CacheMockClient struct {
	mock.Mock
}

func (c *CacheMockClient) Get(key string) []dto.Hotel {
	ret := c.Called(key)
	if ret.Get(0) == nil {
		return nil
	}
	return ret.Get(0).([]dto.Hotel)
}

func (c *CacheMockClient) SetWithExpiration(key string, value []dto.Hotel, expiration int) {
	c.Called(key, value, expiration)
}