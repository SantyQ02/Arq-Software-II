package client

import (
	// "errors"
	"mvc-go/dto"
	// "net/http"

	// log "github.com/sirupsen/logrus"
)

type solrClient struct{}

type solrClientInterface interface {
	SearchHotels(city string) []dto.Hotel
	UpdateHotel(hotelId string) dto.Hotel
	AddHotel(hotelDto dto.Hotel) dto.Hotel
}

var (
	SolrClient solrClientInterface
)

func init() {
	SolrClient = &solrClient{}
}

func (c *solrClient) SearchHotels(city string) []dto.Hotel {
	var hotels []dto.Hotel
	return hotels
}

func (c *solrClient) UpdateHotel(hotelId string) dto.Hotel {
	var hotel dto.Hotel
	return hotel
}

func (c *solrClient) AddHotel(hotelDto dto.Hotel) dto.Hotel {
	var hotel dto.Hotel
	return hotel
}