package service

import (
	"mvc-go/dto"
	"mvc-go/client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"errors"
	"time"
	"fmt"
	// log "github.com/sirupsen/logrus"
	// log "github.com/sirupsen/logrus"
)

func initTestClient2() {
	client.SolrClient = &client.SolrMockClient{}
	client.BusinessClient = &client.BusinessMockClient{}
	client.CacheClient = &client.CacheMockClient{}
}

func TestSearch(t *testing.T) {
	initTestClient2()
	solrMockClient := client.SolrClient.(*client.SolrMockClient)
	businessMockClient := client.BusinessClient.(*client.BusinessMockClient)
	cacheMockClient := client.CacheClient.(*client.CacheMockClient)

	
	city := "CBA"
	hotel_id_1 := uuid.New()
	hotel_id_2 := uuid.New()
	hotel_id_3 := uuid.New()
	
	var hotelsDto = []dto.Hotel{
		dto.Hotel{
			HotelID: hotel_id_1,
		},
		dto.Hotel{
			HotelID: hotel_id_2,
		},
		dto.Hotel{
			HotelID: hotel_id_3,
		},
	}
	
	solrMockClient.On("SearchHotels", city).Return(hotelsDto, nil)
	
	checkInDate, _ := time.Parse("2006-01-02", "2023-10-25")
	checkOutDate, _ := time.Parse("2006-01-02", "2023-11-25")
	
	cacheKey := fmt.Sprintf("%s/%s/%s", city, checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))
	cacheMockClient.On("Get", cacheKey).Return(nil)
	cacheMockClient.On("SetWithExpiration", cacheKey, hotelsDto, 5).Return()

	businessMockClient.On("GetHotelAvailability", hotel_id_1, checkInDate, checkOutDate).Return(dto.BusinessResponse{hotel_id_1, true}, nil)
	businessMockClient.On("GetHotelAvailability", hotel_id_2, checkInDate, checkOutDate).Return(dto.BusinessResponse{hotel_id_2, false}, nil)
	businessMockClient.On("GetHotelAvailability", hotel_id_3, checkInDate, checkOutDate).Return(dto.BusinessResponse{hotel_id_3, true}, nil)

	hotelsRes, err := SearchService.Search(city, checkInDate, checkOutDate)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(hotelsRes))
	assert.Equal(t, true, hotelsRes[0].Available)
	assert.Equal(t, false, hotelsRes[1].Available)
	assert.Equal(t, true, hotelsRes[2].Available)
}

func TestSearchErrorSolrClient(t *testing.T) {
	initTestClient2()
	solrMockClient := client.SolrClient.(*client.SolrMockClient)
	cacheMockClient := client.CacheClient.(*client.CacheMockClient)

	city := "CBA"

	solrMockClient.On("SearchHotels", city).Return([]dto.Hotel{}, errors.New("cualquier error"))

	checkInDate, _ := time.Parse("2006-01-02", "2023-10-25")
	checkOutDate, _ := time.Parse("2006-01-02", "2023-11-25")

	cacheKey := fmt.Sprintf("%s/%s/%s", city, checkInDate.Format("2006-01-02"), checkOutDate.Format("2006-01-02"))
	cacheMockClient.On("Get", cacheKey).Return(nil)

	hotelsRes, err := SearchService.Search(city, checkInDate, checkOutDate)
	assert.NotNil(t, err)
	assert.Equal(t, []dto.Hotel{}, hotelsRes)
	assert.Equal(t, 500, err.Status())

}
