package service

import (
	"mvc-go/dto"
	"mvc-go/client"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initTestClient() {
	client.SolrClient = &client.SolrMockClient{}
	client.HotelClient = &client.HotelMockClient{}
}

func TestAddOrUpdateHotel(t *testing.T) {
	initTestClient()
	solrMockClient := client.SolrClient.(*client.SolrMockClient)
	hotelMockClient := client.HotelClient.(*client.HotelMockClient)
	hotel_id := uuid.New()

	hotelMockClient.On("GetHotel", hotel_id).Return(dto.Hotel{}, nil)
	solrMockClient.On("AddOrUpdateHotel", dto.Hotel{}).Return(nil)

	err := SolrService.AddOrUpdateHotel(hotel_id)

	assert.Nil(t, err)
}

func TestAddOrUpdateHotelErrorHotelClient(t *testing.T) {
	initTestClient()
	hotelMockClient := client.HotelClient.(*client.HotelMockClient)
	hotel_id := uuid.New()

	hotelMockClient.On("GetHotel", hotel_id).Return(dto.Hotel{}, errors.New("Hotel Client Error"))

	err := SolrService.AddOrUpdateHotel(hotel_id)

	assert.NotNil(t, err)
	assert.Equal(t, "Hotel Client Error", err.Error())
}

func TestAddOrUpdateHotelErrorSolrClient(t *testing.T) {
	initTestClient()
	solrMockClient := client.SolrClient.(*client.SolrMockClient)
	hotelMockClient := client.HotelClient.(*client.HotelMockClient)
	hotel_id := uuid.New()

	hotelMockClient.On("GetHotel", hotel_id).Return(dto.Hotel{}, nil)
	solrMockClient.On("AddOrUpdateHotel", dto.Hotel{}).Return(errors.New("Solr Client Error"))

	err := SolrService.AddOrUpdateHotel(hotel_id)

	assert.NotNil(t, err)
	assert.Equal(t, "Solr Client Error", err.Error())
}