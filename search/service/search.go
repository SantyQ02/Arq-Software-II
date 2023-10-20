package service

import (
	// "errors"
	"mvc-go/client"
	"mvc-go/dto"

	e "mvc-go/utils/errors"
	// "github.com/google/uuid"
)

type searchService struct{}

type searchServiceInterface interface {
	Search(city string, checkInDate string, checkOutDate string) ([]dto.Hotel, e.ApiError)
}

var (
	SearchService searchServiceInterface
)

func init() {
	SearchService = &searchService{}
}

func (s *searchService) Search(city string, checkInDate string, checkOutDate string) ([]dto.Hotel, e.ApiError) {

	// fetch hotels
	hotels := client.SolrClient.SearchHotels(city)

	// for each hotel start a go rutine and fetch availability from business micro service

	return hotels, nil
}
