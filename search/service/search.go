package service

import (
	"fmt"
	"mvc-go/client"
	"mvc-go/dto"
	e "mvc-go/utils/errors"

	// "sync"
	"time"
	// "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type searchService struct{}

type searchServiceInterface interface {
	Search(city string, checkInDate time.Time, checkOutDate time.Time) ([]dto.Hotel, e.ApiError)
}

var (
	SearchService searchServiceInterface
)

func init() {
	SearchService = &searchService{}
}

func (s *searchService) Search(city string, checkInDate time.Time, checkOutDate time.Time) ([]dto.Hotel, e.ApiError) {
	cacheKey := fmt.Sprintf("%s/%s/%s", city, checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))

	hotelsCache := client.CacheClient.Get(cacheKey)
	if hotelsCache != nil {
		return hotelsCache, nil
	}

	// fetch hotels
	hotels, err := client.SolrClient.SearchHotels(city)
	if err != nil {
		return []dto.Hotel{}, e.NewInternalServerApiError("Something went wrong searching hotels", err)
	}
	if len(hotels) == 0 {
		return []dto.Hotel{}, nil
	}

	// for each hotel start a go rutine and fetch availability from business micro service
	// var wg sync.WaitGroup
	errCh := make(chan error, len(hotels))

	for i := range hotels {
		// wg.Add(1)
		go func(i int, errCh chan error) {
			// defer wg.Done()
			hotelCopy := &hotels[i]
			businessRes, er := client.BusinessClient.GetHotelAvailability(hotelCopy.HotelID, checkInDate, checkOutDate)
			if er != nil {
				errCh <- er
				hotelCopy.Available = false
				return
			}
			if hotelCopy.HotelID == businessRes.HotelID {
				hotelCopy.Available = businessRes.Available
			} else {
				hotelCopy.Available = false
			}
			errCh <- nil

		}(i, errCh)
	}

	for _ = range hotels {
		err := <-errCh
		if err != nil {
			log.Error(err.Error())
			// return []dto.Hotel{}, e.NewInternalServerApiError("Something went wrong checking availability", err)
		}
	}

	client.CacheClient.SetWithExpiration(cacheKey, hotels, 5)

	return hotels, nil
}
