package businessController

import (
	"bytes"
	"encoding/json"
	"mvc-go/dto"
	businessService "mvc-go/services/business"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func initTestClient() {
	businessService.BusinessService = &businessService.BusinessMockService{}
	gin.SetMode(gin.TestMode)
}

// Tests for CheckAvailability
func TestCheckAvailability(t *testing.T) {
	initTestClient()
	mockBusinessService := businessService.BusinessService.(*businessService.BusinessMockService)

	hotelID := uuid.New()
	checkInDateString := "2023-12-12"
	checkOutDateString := "2023-12-13"
	checkInDate, _ := time.Parse("2006-01-02", checkInDateString)
	checkOutDate, _ := time.Parse("2006-01-02", checkOutDateString)

	mockBusinessService.On("CheckAvailability", hotelID, checkInDate, checkOutDate).Return(true, nil)

	router := gin.Default()
	router.GET("/test/hotel/:hotelID/availability", CheckAvailability)

	req, _ := http.NewRequest("GET", "/test/hotel/"+hotelID.String()+"/availability?checkInDate="+checkInDateString+"&checkOutDate="+checkOutDateString, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), `"hotel_id":"`+hotelID.String()+`"`)
	assert.Contains(t, resp.Body.String(), `"available":true`)
}

func TestCheckAvailabilityErrorInvalidHotelID(t *testing.T) {
	initTestClient()
	router := gin.Default()
	router.GET("/test/hotel/:hotelID/availability", CheckAvailability)

	req, _ := http.NewRequest("GET", "/test/hotel/invalidID/availability?checkInDate=2023-01-01&checkOutDate=2023-01-03", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "hotelID must be a uuid")
}

func TestCheckAvailabilityErrorInvalidCheckInDate(t *testing.T) {
	initTestClient()
	router := gin.Default()
	router.GET("/test/hotel/:hotelID/availability", CheckAvailability)

	hotelID := uuid.New()
	req, _ := http.NewRequest("GET", "/test/hotel/"+hotelID.String()+"/availability?checkInDate=invalidDate&checkOutDate=2023-01-03", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "checkInDate must be a valid value")
}

// Add more test cases for CheckAvailability as needed

// Tests for MapHotel
func TestMapHotel(t *testing.T) {
	initTestClient()
	mockBusinessService := businessService.BusinessService.(*businessService.BusinessMockService)

	hotelMapping := dto.HotelMapping{
		AmadeusID: "ABCDEFG",
		HotelID:   uuid.New(),
	}

	mockBusinessService.On("MapHotel", hotelMapping).Return(hotelMapping, nil)

	router := gin.Default()
	router.POST("/test/hotel/map", MapHotel)

	payload, _ := json.Marshal(hotelMapping)
	req, _ := http.NewRequest("POST", "/test/hotel/map", bytes.NewBuffer(payload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 201, resp.Code)
}

func TestMapHotelErrorInvalidPayload(t *testing.T) {
	initTestClient()

	hotelMapping := dto.HotelMapping{}

	router := gin.Default()
	router.POST("/test/hotel/map", MapHotel)

	payload, _ := json.Marshal(hotelMapping)
	req, _ := http.NewRequest("POST", "/test/hotel/map", bytes.NewBuffer(payload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}