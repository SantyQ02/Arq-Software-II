package hotelService

import (
	"errors"
	hotelClient "mvc-go/clients/hotel"
	"mvc-go/model"
	"mvc-go/dto"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initTestClient() {

	hotelClient.HotelClient = &hotelClient.HotelMockClient{}

}

func TestInsertHotel(t *testing.T){
	initTestClient()

	hotelID := uuid.New()

	hotelModel = model.Hotel{
	HotelID: 	 	hotelID.String(),
	AmadeusID:      "0000",
	Title:         	"Test",
	Description:    "Test desciption",
	PricePerDay:    999,
	CityCode:       "City",
	Photos:         nil,
	Amenities: 		nil,
	Active:         true,
	}

	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("InsertHotel", hotelModel).Return(hotelModel)

	hotelDto = dto.Hotel{
		HotelID: 	 	hotelID,
		AmadeusID:      "0000",
		Title:         	"Test",
		Description:    "Test desciption",
		PricePerDay:    999,
		CityCode:       "City",
		Photos:         nil,
		Amenities: 		nil,
		Active:         true,
		}

	hotel, err := hotelService.InsertHotel(hotelDto)

	assert.Nil(t, err)
	assert.Equal(t, hotelID, hotel.HotelID)
	assert.Equal(t, hotelDto.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelDto.Title, hotel.Title)
	assert.Equal(t, hotelDto.Description, hotel.Description)
	assert.Equal(t, hotelDto.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelDto.CityCode, hotel.CityCode)
	assert.Equal(t, hotelDto.Photos, hotel.Photos)
	assert.Equal(t, hotelDto.Amenities, hotel.Amenities)
	assert.Equal(t, hotelDto.Active, hotel.Active)

}

func TestGetHotelById(t *testing.T){
	initTestClient()

	hotelID := uuid.New()

	hotelModel = model.Hotel{
	HotelID: 	 	hotelID.String(),
	AmadeusID:      "0000",
	Title:         	"Test",
	Description:    "Test desciption",
	PricePerDay:    999,
	CityCode:       "City",
	Photos:         nil,
	Amenities: 		nil,
	Active:         true,
	}

	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("GetHotelByID", hotelID).Return(hotelModel)

	hotel,err := hotelService.GetHotelByID(hotelID)

	assert.Nil(t, err)
	assert.Equal(t, hotelID, hotel.HotelID)
	assert.Equal(t, hotelModel.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelModel.Title, hotel.Title)
	assert.Equal(t, hotelModel.Description, hotel.Description)
	assert.Equal(t, hotelModel.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelModel.CityCode, hotel.CityCode)
	assert.Equal(t, hotelModel.Photos, hotel.Photos)
	assert.Equal(t, hotelModel.Amenities, hotel.Amenities)
	assert.Equal(t, hotelModel.Active, hotel.Active)

}

func TestUpdateHotel(t *testing.T){
	initTestClient()

	hotelID := uuid.New()

	hotelModel = model.Hotel{
	HotelID: 	 	hotelID.String(),
	AmadeusID:      "0000",
	Title:         	"Test",
	Description:    "Test desciption",
	PricePerDay:    888,
	CityCode:       "City",
	Photos:         nil,
	Amenities: 		nil,
	Active:         true,
	}

	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("UpdateHotel", hotelModel).Return(hotelModel)

	hotelDto = dto.Hotel{
		HotelID: 	 	hotelID,
		AmadeusID:      "0000",
		Title:         	"Test",
		Description:    "Test desciption",
		PricePerDay:    888,
		CityCode:       "City",
		Photos:         nil,
		Amenities: 		nil,
		Active:         true,
		}

	hotel,err := hotelService.UpdateHotel(hotelDto)

	assert.Nil(t, err)
	assert.Equal(t, hotelID, hotel.HotelID)
	assert.Equal(t, hotelDto.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelDto.Title, hotel.Title)
	assert.Equal(t, hotelDto.Description, hotel.Description)
	assert.Equal(t, hotelDto.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelDto.CityCode, hotel.CityCode)
	assert.Equal(t, hotelDto.Photos, hotel.Photos)
	assert.Equal(t, hotelDto.Amenities, hotel.Amenities)
	assert.Equal(t, hotelDto.Active, hotel.Active)
}

func TestDeleteHotel(t *testing.T){
	initTestClient()

	hotelID := uuid.New()

	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("DeleteHotel", hotelID).Return(nil)

	err := hotelService.DeleteHotel(hotelID)

	assert.Nil(t, err)
}

func TestInsertPhoto(t *testing.T){
	initTestClient()

	hotelID := uuid.New()
	photoID := uuid.New()
	photoDto := dto.Photo{
                PhotoID: photoID,
                Url: "test/url"
            }
	photosDto := dto.Photos{photoDto}

	hotelModel = model.Hotel{
	HotelID: 	 	hotelID.String(),
	AmadeusID:      "0000",
	Title:         	"Test",
	Description:    "Test desciption",
	PricePerDay:    888,
	CityCode:       "City",
	Photos:         photosDto,
	Amenities: 		nil,
	Active:         true,
	}
        
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("UpdateHotel", hotelModel).Return(hotelModel)

	photo, err := hotelService.InsertPhoto(photoDto, hotelID)

	assert.Nil(t, err)
	assert.Equal(t, photo.PhotoID, photoID)
	assert.Equal(t, photo.Url, photoDto.Url)

}

