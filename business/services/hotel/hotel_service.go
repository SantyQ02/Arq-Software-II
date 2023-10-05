package hotelService

import (
	"errors"
	hotelClient "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"

	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
)

type hotelService struct{}

type hotelServiceInterface interface {
	InsertHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError)
	UpdateHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError)
	DeleteHotel(id uuid.UUID) e.ApiError
	GetHotels() ([]dto.Hotel, e.ApiError)
	GetHotelById(id uuid.UUID) (dto.Hotel, e.ApiError)
	GetAvailableRooms(booking dto.CheckAvailability) (float64, e.ApiError)
	GetAvailableHotels(booking dto.CheckAvailability) (dto.Hotels, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) InsertHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError) {
	hotel := model.Hotel{
		Title:       hotelDto.Title,
		Description: hotelDto.Description,
		Rooms:       hotelDto.Rooms,
		PricePerDay: hotelDto.PricePerDay,
		Active:      true,
	}

	hotel = hotelClient.HotelClient.InsertHotel(hotel)
	if hotel.HotelID == uuid.Nil {
		return dto.Hotel{}, e.NewInternalServerApiError("Error trying insert new hotel", errors.New(""))
	}

	hotelDto.HotelID = hotel.HotelID

	return hotelDto, nil
}

func (s *hotelService) UpdateHotel(hotelDto dto.Hotel) (dto.Hotel, e.ApiError) {

	hotel := model.Hotel{
		HotelID:     hotelDto.HotelID,
		Title:       hotelDto.Title,
		Description: hotelDto.Description,
		Rooms:       hotelDto.Rooms,
		PricePerDay: hotelDto.PricePerDay,
		Active:      hotelDto.Active,
	}

	hotel = hotelClient.HotelClient.UpdateHotel(hotel)
	if hotel.HotelID == uuid.Nil {
		return dto.Hotel{}, e.NewInternalServerApiError("Error trying update hotel", errors.New(""))
	}

	hotelDto.HotelID = hotel.HotelID

	return hotelDto, nil
}

func (s *hotelService) DeleteHotel(id uuid.UUID) e.ApiError {
	idString := id.String()

	err := hotelClient.HotelClient.DeleteHotel(idString)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong deleting hotel", nil)
	}

	return nil

}

func (s *hotelService) GetHotels() ([]dto.Hotel, e.ApiError) {
	hotels := hotelClient.HotelClient.GetHotels()
	if len(hotels) == 0 {
		return []dto.Hotel{}, e.NewInternalServerApiError("Error getting hotels from database", errors.New("error in database"))
	}

	var hotelsDto []dto.Hotel

	for _, hotel := range hotels {
		var hotelDto dto.Hotel
		hotelDto.HotelID = hotel.HotelID
		hotelDto.Title = hotel.Title
		hotelDto.Description = hotel.Description
		hotelDto.Rooms = hotel.Rooms
		hotelDto.PricePerDay = hotel.PricePerDay
		hotelDto.Active = hotel.Active
		for _, photo := range hotel.Photos {
			var dtoPhoto dto.Photo

			dtoPhoto.PhotoID = photo.PhotoID
			dtoPhoto.Url = photo.Url
			dtoPhoto.HotelID = photo.HotelID

			hotelDto.Photos = append(hotelDto.Photos, dtoPhoto)
		}
		for _, amenity := range hotel.Amenities {
			var dtoAmenity dto.Amenitie

			dtoAmenity.AmenitieID = amenity.AmenitieID
			dtoAmenity.Title = amenity.Title

			hotelDto.Amenities = append(hotelDto.Amenities, dtoAmenity)
		}

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) GetHotelById(id uuid.UUID) (dto.Hotel, e.ApiError) {
	idString := id.String()

	hotel := hotelClient.HotelClient.GetHotelById(idString)

	if hotel.HotelID == uuid.Nil {
		return dto.Hotel{}, e.NewNotFoundApiError("Hotel not found")
	}

	hotelDto := dto.Hotel{
		HotelID:     hotel.HotelID,
		Title:       hotel.Title,
		Description: hotel.Description,
		Rooms:       hotel.Rooms,
		PricePerDay: hotel.PricePerDay,
		Active:      hotel.Active,
	}
	for _, photo := range hotel.Photos {
		var dtoPhoto dto.Photo

		dtoPhoto.PhotoID = photo.PhotoID
		dtoPhoto.Url = photo.Url
		dtoPhoto.HotelID = photo.HotelID

		hotelDto.Photos = append(hotelDto.Photos, dtoPhoto)
	}
	for _, amenity := range hotel.Amenities {
		var dtoAmenity dto.Amenitie

		dtoAmenity.AmenitieID = amenity.AmenitieID
		dtoAmenity.Title = amenity.Title

		hotelDto.Amenities = append(hotelDto.Amenities, dtoAmenity)
	}

	return hotelDto, nil
}

func (s *hotelService) GetAvailableRooms(booking dto.CheckAvailability) (float64, e.ApiError) {
	if booking.DateIn.Before(time.Now()) {
		return 0, e.NewBadRequestApiError("You should not have a DateIn earlier than the current date")
	}

	if booking.DateIn.After(booking.DateOut) || booking.DateIn.Equal(booking.DateOut) {
		return 0, e.NewBadRequestApiError("You should not have a DateIn greater or equal than the DateOut")
	}

	availableRooms := hotelClient.HotelClient.GetAvailableRooms(booking)

	return availableRooms, nil
}

func (s *hotelService) GetAvailableHotels(booking dto.CheckAvailability) (dto.Hotels, e.ApiError) {
	hotels := hotelClient.HotelClient.GetAvailableHotels(booking)

	if booking.DateIn.Before(time.Now()) {
		return []dto.Hotel{}, e.NewBadRequestApiError("You should not have a DateIn earlier than the current date")
	}

	if booking.DateIn.After(booking.DateOut) || booking.DateIn.Equal(booking.DateOut) {
		return []dto.Hotel{}, e.NewBadRequestApiError("You should not have a DateIn greater or equal than the DateOut")
	}

	if booking.Rooms == 0 {
		return []dto.Hotel{}, e.NewBadRequestApiError("Error getting hotels from database: You must provide a 'rooms' value")
	}

	var hotelsDto []dto.Hotel

	for _, hotel := range hotels {
		var hotelDto dto.Hotel
		hotelDto.HotelID = hotel.HotelID
		hotelDto.Title = hotel.Title
		hotelDto.Description = hotel.Description
		hotelDto.Rooms = hotel.Rooms
		hotelDto.PricePerDay = hotel.PricePerDay
		hotelDto.Active = hotel.Active
		for _, photo := range hotel.Photos {
			var dtoPhoto dto.Photo

			dtoPhoto.PhotoID = photo.PhotoID
			dtoPhoto.Url = photo.Url
			dtoPhoto.HotelID = photo.HotelID

			hotelDto.Photos = append(hotelDto.Photos, dtoPhoto)
		}
		for _, amenity := range hotel.Amenities {
			var dtoAmenity dto.Amenitie

			dtoAmenity.AmenitieID = amenity.AmenitieID
			dtoAmenity.Title = amenity.Title

			hotelDto.Amenities = append(hotelDto.Amenities, dtoAmenity)
		}

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}
