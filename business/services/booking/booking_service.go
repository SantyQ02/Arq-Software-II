package bookingService

import (
	"errors"
	bookingClient "mvc-go/clients/booking"
	hotelClient "mvc-go/clients/hotel"
	"mvc-go/dto"
	"mvc-go/model"
	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
)

type bookingService struct{}

type bookingServiceInterface interface {
	CreateBooking(bookingDto dto.Booking) (dto.Booking, e.ApiError)
	GetBookingById(booking_id uuid.UUID) (dto.Booking, e.ApiError)
	GetBookings() (dto.Bookings, e.ApiError)
	SearchBookings(hotel string, user string, dateIn time.Time, dateOut time.Time) (dto.Bookings, e.ApiError)
	GetBookingsByUserId(id uuid.UUID) (dto.Bookings, e.ApiError)
	DeleteBooking(id uuid.UUID) e.ApiError
	SetActiveBooking(dto.Booking) (dto.Booking, e.ApiError)
}

var (
	BookingService bookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}

func (s *bookingService) CreateBooking(bookingDto dto.Booking) (dto.Booking, e.ApiError) {
	booking := model.Booking{
		Rooms:   bookingDto.Rooms,
		Total:   bookingDto.Total,
		DateIn:  bookingDto.DateIn,
		DateOut: bookingDto.DateOut,
		UserID:  bookingDto.UserID,
		HotelID: bookingDto.HotelID,
		Active:  true,
	}

	bookingData := dto.CheckAvailability{
		HotelID: booking.HotelID,
		DateIn:  booking.DateIn,
		DateOut: booking.DateOut,
	}

	if booking.Total <= 0 {
		return dto.Booking{}, e.NewBadRequestApiError("You cannot have a zero or negative amount for Total value")
	}

	if booking.DateIn.Before(time.Now()) {
		return dto.Booking{}, e.NewBadRequestApiError("You should not have a DateIn earlier than the current date")
	}

	if booking.DateIn.After(booking.DateOut) || booking.DateIn.Equal(booking.DateOut) {
		return dto.Booking{}, e.NewBadRequestApiError("You should not have a DateIn greater or equal than the DateOut")
	}

	availableRooms := hotelClient.HotelClient.GetAvailableRooms(bookingData)
	if booking.Rooms > uint(availableRooms) {
		return dto.Booking{}, e.NewBadRequestApiError("You cannot book more rooms than the ones currently available")
	}

	booking = bookingClient.BookingClient.InsertBooking(booking)
	if booking.BookingID == uuid.Nil {
		return dto.Booking{}, e.NewInternalServerApiError("Error trying to create new booking", errors.New(""))
	}

	bookingDto.BookingID = booking.BookingID

	return bookingDto, nil
}

func (s *bookingService) GetBookingById(booking_id uuid.UUID) (dto.Booking, e.ApiError) {

	booking := bookingClient.BookingClient.GetBookingById(booking_id.String())
	if booking.BookingID == uuid.Nil {
		return dto.Booking{}, e.NewInternalServerApiError("Error getting bookings from database", errors.New("error in database"))
	}

	bookingDto := dto.Booking{
		BookingID: booking.BookingID,
		Total:     booking.Total,
		Rooms:     booking.Rooms,
		UserID:    booking.UserID,
		HotelID:   booking.HotelID,
		DateIn:    booking.DateIn,
		DateOut:   booking.DateOut,
		Active:    booking.Active,
	}

	return bookingDto, nil
}

func (s *bookingService) GetBookings() (dto.Bookings, e.ApiError) {
	bookings := bookingClient.BookingClient.GetBookings()
	if len(bookings) == 0 {
		return dto.Bookings{}, e.NewInternalServerApiError("Error getting bookings from database", errors.New("error in database"))
	}

	var bookingsDto dto.Bookings

	for _, booking := range bookings {
		var bookingDto dto.Booking
		bookingDto.BookingID = booking.BookingID
		bookingDto.UserID = booking.UserID
		bookingDto.Rooms = booking.Rooms
		bookingDto.Total = booking.Total
		bookingDto.DateIn = booking.DateIn
		bookingDto.DateOut = booking.DateOut
		bookingDto.HotelID = booking.HotelID
		bookingDto.Active = booking.Active

		bookingsDto = append(bookingsDto, bookingDto)
	}

	return bookingsDto, nil
}

func (s *bookingService) GetBookingsByUserId(id uuid.UUID) (dto.Bookings, e.ApiError) {
	idString := id.String()
	bookings := bookingClient.BookingClient.GetBookingsByUserId(idString)
	if len(bookings) == 0 {
		return dto.Bookings{}, e.NewNotFoundApiError("Bookings not found")

	}

	var bookingsDto dto.Bookings

	for _, booking := range bookings {
		var bookingDto dto.Booking
		bookingDto.BookingID = booking.BookingID
		bookingDto.UserID = booking.UserID
		bookingDto.Rooms = booking.Rooms
		bookingDto.Total = booking.Total
		bookingDto.DateIn = booking.DateIn
		bookingDto.DateOut = booking.DateOut
		bookingDto.HotelID = booking.HotelID
		bookingDto.Active = booking.Active

		bookingsDto = append(bookingsDto, bookingDto)
	}

	return bookingsDto, nil
}

func (s *bookingService) SearchBookings(hotel string, user string, dateIn time.Time, dateOut time.Time) (dto.Bookings, e.ApiError) {
	var bookings model.Bookings
	// if hotel == "" {
	// 	bookings = bookingClient.BookingClient.SearchBookingsByDates(dateIn, dateOut)
	// } else {

	bookings = bookingClient.BookingClient.SearchBookingsByDatesAndHotelAndUser(hotel, user, dateIn, dateOut)
	// }
	// if len(bookings) == 0 {
	// 	return dto.Bookings{}, e.NewInternalServerApiError("Error getting bookings from database", errors.New("Error in database"))
	// }

	var bookingsDto dto.Bookings

	for _, booking := range bookings {
		var bookingDto dto.Booking
		bookingDto.BookingID = booking.BookingID
		bookingDto.UserID = booking.UserID
		bookingDto.Rooms = booking.Rooms
		bookingDto.Total = booking.Total
		bookingDto.DateIn = booking.DateIn
		bookingDto.DateOut = booking.DateOut
		bookingDto.HotelID = booking.HotelID
		bookingDto.Active = booking.Active

		bookingsDto = append(bookingsDto, bookingDto)
	}

	return bookingsDto, nil
}

func (s *bookingService) DeleteBooking(id uuid.UUID) e.ApiError {
	idString := id.String()

	err := bookingClient.BookingClient.DeleteBooking(idString)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong deleting booking", nil)
	}

	return nil

}

func (s *bookingService) SetActiveBooking(bookingDto dto.Booking) (dto.Booking, e.ApiError) {

	if bookingDto.DateIn.Before(time.Now()) {
		return dto.Booking{}, e.NewBadRequestApiError("You cannot modify bookings of past dates")
	}

	booking := model.Booking{
		BookingID: bookingDto.BookingID,
		HotelID:   bookingDto.HotelID,
		UserID:    bookingDto.UserID,
		Rooms:     bookingDto.Rooms,
		Total:     bookingDto.Total,
		DateIn:    bookingDto.DateIn,
		DateOut:   bookingDto.DateOut,
		Active:    bookingDto.Active,
	}

	booking = bookingClient.BookingClient.UpdateBooking(booking)
	if booking.BookingID == uuid.Nil {
		return dto.Booking{}, e.NewInternalServerApiError("Something went wrong deleting booking", nil)
	}

	return bookingDto, nil

}