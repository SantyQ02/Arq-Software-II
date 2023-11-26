package businessService

import (
	"errors"
	businessClient "mvc-go/clients/business"
	userClient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"
	"time"
	e "mvc-go/utils/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "os"
)

func initTestClient() {
	businessClient.BusinessClient = &businessClient.BusinessMockClient{}
	userClient.UserClient = &userClient.UserMockClient{}
}

// CheckAvailability
func TestCheckAvailability(t *testing.T) {
	initTestClient()
	mockClient := businessClient.BusinessClient.(*businessClient.BusinessMockClient)
	hotelID := uuid.New()
	checkInDate := time.Now().UTC()
	checkOutDate := checkInDate.Add(24 * time.Hour).UTC()

	mockClient.On("GetAmadeusIDByHotelID", hotelID).Return("amadeus123")
	mockClient.On("GetAmadeusAvailability", "amadeus123", checkInDate, checkOutDate).Return(true, nil)

	availability, err := BusinessService.CheckAvailability(hotelID, checkInDate, checkOutDate)

	assert.True(t, availability)
	assert.Nil(t, err)
}

func TestCheckAvailabilityErrorHotelNotFound(t *testing.T) {
	initTestClient()
	mockClient := businessClient.BusinessClient.(*businessClient.BusinessMockClient)
	hotelID := uuid.New()
	checkInDate := time.Now().UTC()
	checkOutDate := checkInDate.Add(24 * time.Hour).UTC()

	mockClient.On("GetAmadeusIDByHotelID", hotelID).Return("")

	availability, err := BusinessService.CheckAvailability(hotelID, checkInDate, checkOutDate)

	assert.False(t, availability)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "Hotel not found!", err.Message())
}

func TestCheckAvailabilityErrorAmadeusAPI(t *testing.T) {
	initTestClient()
	mockClient := businessClient.BusinessClient.(*businessClient.BusinessMockClient)
	hotelID := uuid.New()
	checkInDate := time.Now().UTC()
	checkOutDate := checkInDate.Add(24 * time.Hour).UTC()

	mockClient.On("GetAmadeusIDByHotelID", hotelID).Return("amadeus123")
	mockClient.On("GetAmadeusAvailability", "amadeus123", checkInDate, checkOutDate).Return(false, e.NewInternalServerApiError("500", errors.New("500")))

	availability, err := BusinessService.CheckAvailability(hotelID, checkInDate, checkOutDate)

	assert.False(t, availability)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Status())
}

// MapHotel
func TestMapHotel(t *testing.T) {
	initTestClient()
	mockClient := businessClient.BusinessClient.(*businessClient.BusinessMockClient)
	hotelMappingDto := dto.HotelMapping{
		HotelID:   uuid.New(),
		AmadeusID: "amadeus123",
	}

	mockClient.On("InsertHotelMapping", mock.Anything).Return(nil)

	mappedHotel, err := BusinessService.MapHotel(hotelMappingDto)

	assert.Equal(t, hotelMappingDto, mappedHotel)
	assert.Nil(t, err)
}

func TestMapHotelErrorAlreadyMapped(t *testing.T) {
	initTestClient()
	mockClient := businessClient.BusinessClient.(*businessClient.BusinessMockClient)
	hotelMappingDto := dto.HotelMapping{
		HotelID:   uuid.New(),
		AmadeusID: "amadeus123",
	}

	mockClient.On("InsertHotelMapping", mock.Anything).Return(errors.New("400"))

	mappedHotel, err := BusinessService.MapHotel(hotelMappingDto)

	assert.Equal(t, dto.HotelMapping{}, mappedHotel)
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
}

// CheckAdmin
func TestCheckAdmin(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	userID := uuid.New()

	mockClient.On("GetUserById", userID.String()).Return(model.User{UserID: userID, Role: "admin"})

	isAdmin, err := BusinessService.CheckAdmin(userID)

	assert.True(t, isAdmin)
	assert.Nil(t, err)
}

func TestCheckAdminErrorUserNotFound(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	userID := uuid.New()

	mockClient.On("GetUserById", userID.String()).Return(model.User{})

	isAdmin, err := BusinessService.CheckAdmin(userID)

	assert.False(t, isAdmin)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "User not found", err.Message())
}

func TestCheckAdminErrorNotAdmin(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	userID := uuid.New()

	mockClient.On("GetUserById", userID.String()).Return(model.User{UserID: userID, Role: "user"})

	isAdmin, err := BusinessService.CheckAdmin(userID)

	assert.False(t, isAdmin)
	assert.Nil(t, err)
}
