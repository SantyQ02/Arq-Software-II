package service

import (
	"mvc-go/client"
	"mvc-go/dto"

	e "mvc-go/utils/errors"
	"sync"
	"time"
	// "github.com/google/uuid"
	// log "github.com/sirupsen/logrus"
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

	// fetch hotels
	hotels, err := client.SolrClient.SearchHotels(city)
	if err != nil {
		return []dto.Hotel{}, e.NewInternalServerApiError("Something went wrong searching hotels", err)
	}

	// for each hotel start a go rutine and fetch availability from business micro service
	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for i := range hotels {
		wg.Add(1)
		go func (i int, errCh chan error) {
			defer wg.Done()
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

		}(i, errCh)
	}

	wg.Wait()
	close(errCh)
	
	select {
    case err := <-errCh:
        if err != nil {
            return []dto.Hotel{}, e.NewInternalServerApiError("Something went wrong checking availability", err)
        }
    default:
        // No se produjo un error diferente de nil, continúa con la lógica
    }


	return hotels, nil
}
