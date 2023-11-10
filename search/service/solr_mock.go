package service

import (
	"github.com/google/uuid"

	"github.com/stretchr/testify/mock"
)

type SolrMockService struct {
	mock.Mock
}

func (c *SolrMockService) AddOrUpdateHotel(hotel_id uuid.UUID) error {
	ret := c.Called(hotel_id)
	return ret.Error(0)
}