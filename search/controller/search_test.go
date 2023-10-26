package controller

import (
	"mvc-go/dto"
	"mvc-go/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"io/ioutil"
	e "mvc-go/utils/errors"
	"errors"
)

func initTestService() {
	service.SearchService = &service.SearchMockService{}
}

type BodyRes struct {
	Hotels []dto.Hotel `json:"hotels"`
}

type ErrorRes struct {
	Error string `json:"error"`
}

func TestSearch(t *testing.T) {
	initTestService()
	searchMockService := service.SearchService.(*service.SearchMockService)

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

	checkInDate, _ := time.Parse("2006-01-02", "2023-10-25")
	checkOutDate, _ := time.Parse("2006-01-02", "2023-11-25")

	searchMockService.On("Search", city, checkInDate, checkOutDate).Return(hotelsDto, nil)

	router := gin.Default()
	router.GET("/test/search", Search)

	req, _ := http.NewRequest("GET", "/test/search?city=CBA&check_in_date=2023-10-25&check_out_date=2023-11-25", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	bytes, _ := ioutil.ReadAll(resp.Body)

	var body BodyRes
	json.Unmarshal(bytes, &body)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, hotelsDto, body.Hotels)
}

func TestSearchErrorService(t *testing.T) {
	initTestService()
	searchMockService := service.SearchService.(*service.SearchMockService)

	city := "CBA"

	checkInDate, _ := time.Parse("2006-01-02", "2023-10-25")
	checkOutDate, _ := time.Parse("2006-01-02", "2023-11-25")

	searchMockService.On("Search", city, checkInDate, checkOutDate).Return([]dto.Hotel{}, e.NewInternalServerApiError("Something went wrong searching hotels", errors.New("")))

	router := gin.Default()
	router.GET("/test/search", Search)

	req, _ := http.NewRequest("GET", "/test/search?city=CBA&check_in_date=2023-10-25&check_out_date=2023-11-25", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	bytes, _ := ioutil.ReadAll(resp.Body)

	var body ErrorRes
	json.Unmarshal(bytes, &body)

	assert.Equal(t, 500, resp.Code)
	assert.Equal(t, "Something went wrong searching hotels", body.Error)
}

func TestSearchBadRequestCityValue(t *testing.T) {
	initTestService()

	router := gin.Default()
	router.GET("/test/search", Search)

	req, _ := http.NewRequest("GET", "/test/search?city=&check_in_date=2023-10-25&check_out_date=2023-11-25", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	bytes, _ := ioutil.ReadAll(resp.Body)

	var body ErrorRes
	json.Unmarshal(bytes, &body)

	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, "city must be a valid value", body.Error)
}

func TestSearchBadRequestDateValue(t *testing.T) {
	initTestService()

	router := gin.Default()
	router.GET("/test/search", Search)

	req, _ := http.NewRequest("GET", "/test/search?city=CBA&check_in_date=2023-10-25ff&check_out_date=2023-11-25ff", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	bytes, _ := ioutil.ReadAll(resp.Body)

	var body ErrorRes
	json.Unmarshal(bytes, &body)

	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, "check_in_date must be a valid value", body.Error)
}