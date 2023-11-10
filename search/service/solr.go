package service

import (
	// "errors"
	"mvc-go/client"
	// "mvc-go/dto"

	// e "mvc-go/utils/errors"
	"github.com/google/uuid"
)

type solrService struct{}

type solrServiceInterface interface {
	AddOrUpdateHotel(hotel_id uuid.UUID) error
}

var (
	SolrService solrServiceInterface
)

func init() {
	SolrService = &solrService{}
}

func (s *solrService) AddOrUpdateHotel(hotel_id uuid.UUID) error {

	hotel, err := client.HotelClient.GetHotel(hotel_id)
	if err != nil {
		return err
	}

	err = client.SolrClient.AddOrUpdateHotel(hotel)
	if err != nil {
		return err
	}

	return nil
}
