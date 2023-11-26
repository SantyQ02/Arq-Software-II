package hotelService

import (
	"errors"
	hotelClient "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initTestClient() {
	HotelService = &hotelService{}
	hotelClient.HotelClient = &hotelClient.HotelMockClient{}

}

func TestInsertHotel(t *testing.T){
	initTestClient()

	hotelModel := model.Hotel{
	HotelID: 		uuid.New().String(),
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
	mockHotelClient.On("HotelMapping", hotelModel).Return(nil)
	mockHotelClient.On("InsertHotel", mock.Anything).Return(hotelModel)

	hotelDto := dto.Hotel{
		AmadeusID:      "0000",
		Title:         	"Test",
		Description:    "Test desciption",
		PricePerDay:    999,
		CityCode:       "City",
		Photos:         nil,
		Amenities: 		nil,
		Active:         true,
		}

	hotel, err := HotelService.InsertHotel(hotelDto)

	assert.Nil(t, err)
	assert.NotEqual(t, "", hotel.HotelID)
	assert.Equal(t, hotelDto.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelDto.Title, hotel.Title)
	assert.Equal(t, hotelDto.Description, hotel.Description)
	assert.Equal(t, hotelDto.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelDto.CityCode, hotel.CityCode)
	assert.Equal(t, hotelDto.Photos, hotel.Photos)
	assert.Equal(t, hotelDto.Amenities, hotel.Amenities)
	assert.Equal(t, hotelDto.Active, hotel.Active)

}

func TestInsertHotel_Failure(t *testing.T) {
	initTestClient()

	hotelID := uuid.New()

	hotelModel := model.Hotel{
		HotelID:     hotelID.String(),
		AmadeusID:   "0000",
		Title:       "Test",
		Description: "Test description",
		PricePerDay: 999,
		CityCode:    "City",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("HotelMapping", hotelModel).Return(nil)
	mockHotelClient.On("InsertHotel", mock.Anything).Return(model.Hotel{}) // Simulating a failure by returning an empty hotel model

	hotelDto := dto.Hotel{
		AmadeusID:   "0000",
		Title:       "Test",
		Description: "Test description",
		PricePerDay: 999,
		CityCode:    "City",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	hotel, err := HotelService.InsertHotel(hotelDto)

	assert.NotNil(t, err) // Expecting an error
	assert.Equal(t, dto.Hotel{}, hotel) // Expecting an empty hotel object
}


func TestGetHotelById(t *testing.T){
	initTestClient()

	hotelID := uuid.New()

	hotelModel := model.Hotel{
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
	mockHotelClient.On("GetHotelById", hotelID.String()).Return(hotelModel)

	hotel,err := HotelService.GetHotelById(hotelID)

	assert.Nil(t, err)
	assert.Equal(t, hotelID, hotel.HotelID)
	assert.Equal(t, hotelModel.AmadeusID, hotel.AmadeusID)
	assert.Equal(t, hotelModel.Title, hotel.Title)
	assert.Equal(t, hotelModel.Description, hotel.Description)
	assert.Equal(t, hotelModel.PricePerDay, hotel.PricePerDay)
	assert.Equal(t, hotelModel.CityCode, hotel.CityCode)
	assert.Nil(t, hotel.Photos)
	assert.Nil(t, hotel.Amenities)
	assert.Equal(t, hotelModel.Active, hotel.Active)

}

func TestGetHotelById_NotFound(t *testing.T) {
	initTestClient()

	hotelID := uuid.New()

	// Mocking a scenario where GetHotelByID returns an empty model.Hotel
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("GetHotelById", hotelID.String()).Return(model.Hotel{})

	hotel, err := HotelService.GetHotelById(hotelID)

	assert.NotNil(t, err) // Expecting an error indicating that the hotel was not found
	assert.Equal(t, dto.Hotel{}, hotel) // Expecting an empty hotel object
}

func TestGetHotels_Success(t *testing.T) {
	initTestClient()

	// Mocking a scenario where GetHotels returns a list of hotels
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	hotels := model.Hotels{}

	hotel1 := model.Hotel{
		HotelID:     "1",
		AmadeusID:   "0001",
		Title:       "Hotel 1",
		Description: "Description 1",
		PricePerDay: 100,
		CityCode:    "City1",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	hotel2 := model.Hotel{
		HotelID:     "2",
		AmadeusID:   "0002",
		Title:       "Hotel 2",
		Description: "Description 2",
		PricePerDay: 150,
		CityCode:    "City2",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	hotels = append(hotels, hotel1, hotel2)

	mockHotelClient.On("GetHotels").Return(hotels)

	hotelsDto, err := HotelService.GetHotels()

	assert.Nil(t, err)
	assert.NotNil(t, hotelsDto)
	assert.Equal(t, 2, len(hotelsDto)) // Assuming two hotels are returned
}

func TestGetHotels_NoHotelsFound(t *testing.T) {
	initTestClient()

	// Mocking a scenario where GetHotels returns an empty list
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("GetHotels").Return(model.Hotels{})

	hotelsDto, err := HotelService.GetHotels()

	assert.NotNil(t, err) // Expecting an error indicating no hotels found
	assert.Equal(t, dto.Hotels{}, hotelsDto) // Expecting an empty hotelsDto object
}



func TestUpdateHotel(t *testing.T){
	initTestClient()

	hotelID := uuid.New()

	hotelModel := model.Hotel{
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

	hotelDto := dto.Hotel{
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

	hotel,err := HotelService.UpdateHotel(hotelDto)

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

func TestUpdateHotel_Failure(t *testing.T) {
	initTestClient()

	hotelID := uuid.New()

	hotelModel := model.Hotel{
		HotelID:     hotelID.String(),
		AmadeusID:   "0000",
		Title:       "Test",
		Description: "Test description",
		PricePerDay: 888,
		CityCode:    "City",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	// Mocking a scenario where UpdateHotel returns an empty model.Hotel
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("UpdateHotel", hotelModel).Return(model.Hotel{})

	hotelDto := dto.Hotel{
		HotelID:     hotelID,
		AmadeusID:   "0000",
		Title:       "Test",
		Description: "Test description",
		PricePerDay: 888,
		CityCode:    "City",
		Photos:      nil,
		Amenities:   nil,
		Active:      true,
	}

	hotel, err := HotelService.UpdateHotel(hotelDto)

	assert.NotNil(t, err) // Expecting an error indicating the update failure
	assert.Equal(t, dto.Hotel{}, hotel) // Expecting an empty hotel object
}


func TestDeleteHotel(t *testing.T){
	initTestClient()

	hotelID := uuid.New()

	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("DeleteHotel", hotelID.String(),).Return(nil)

	err := HotelService.DeleteHotel(hotelID)

	assert.Nil(t, err)
}

func TestDeleteHotel_Failure(t *testing.T) {
	initTestClient()

	hotelID := uuid.New()

	// Mocking a scenario where DeleteHotel returns an error
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("DeleteHotel", hotelID.String()).Return(errors.New("deletion failed"))

	err := HotelService.DeleteHotel(hotelID)

	assert.NotNil(t, err) // Expecting an error indicating the deletion failure
}


func TestInsertPhoto(t *testing.T){
	initTestClient()

	hotelID := uuid.New()
	photoDto := dto.Photo{
                Url: "test/url",
            }
	photoModel := model.Photo{
		Url:     photoDto.Url,
	}
	photosModel := model.Photos{photoModel}

	hotelModel1 := model.Hotel{
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
	
	hotelModel2 := model.Hotel{
	HotelID: 	 	hotelID.String(),
	AmadeusID:      "0000",
	Title:         	"Test",
	Description:    "Test desciption",
	PricePerDay:    888,
	CityCode:       "City",
	Photos:         photosModel,
	Amenities: 		nil,
	Active:         true,
	}
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("GetHotelById", hotelID.String()).Return(hotelModel1)
	mockHotelClient.On("UpdateHotel",  mock.AnythingOfType("model.Hotel")).Return(hotelModel2)

	photo, err := HotelService.UploadPhoto(photoDto, hotelID)

	assert.Nil(t, err)
	assert.NotNil(t, photo.PhotoID)
	assert.Equal(t, photo.Url, photoDto.Url)

}

func TestUploadPhoto_UpdateFailure(t *testing.T) {
	initTestClient()

	hotelID := uuid.New()
	photoDto := dto.Photo{
		Url:     "test/url",
	}

	hotelModel := model.Hotel{
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
	
	// Mocking a scenario where UpdateHotel returns an empty model.Hotel
	mockHotelClient := hotelClient.HotelClient.(*hotelClient.HotelMockClient)
	mockHotelClient.On("GetHotelById", hotelID.String()).Return(hotelModel)
	mockHotelClient.On("UpdateHotel",  mock.AnythingOfType("model.Hotel")).Return(model.Hotel{})

	photo, err := HotelService.UploadPhoto(photoDto, hotelID)

	assert.NotNil(t, err) // Expecting an error indicating the update failure
	assert.Equal(t, dto.Photo{}, photo) // Expecting an empty photo object
}


