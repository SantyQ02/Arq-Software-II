package hotelClient

import (
	"errors"
	"mvc-go/dto"
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type hotelClient struct{}

type hotelClientInterface interface {
	GetHotelById(id string) model.Hotel
	GetHotels() model.Hotels
	InsertHotel(hotel model.Hotel) model.Hotel
	UpdateHotel(hotel model.Hotel) model.Hotel
	DeleteHotel(id string) error
	GetAvailableRooms(booking dto.CheckAvailability) float64
	GetAvailableHotels(booking dto.CheckAvailability) model.Hotels
}

var (
	HotelClient hotelClientInterface
)

func init() {
	HotelClient = &hotelClient{}
}

var Db *gorm.DB

func (c *hotelClient) GetHotelById(id string) model.Hotel {
	var hotel model.Hotel

	Db.Preload("Photos").Preload("Amenities").First(&hotel, "hotel_id = ?", id)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func (c *hotelClient) GetHotels() model.Hotels {
	var hotels model.Hotels
	result := Db.Preload("Photos").Preload("Amenities").Find(&hotels)
	if result.Error != nil {
		log.Error("")
		return model.Hotels{}
	}
	log.Debug("Hotels: ", hotels)

	return hotels
}

func (c *hotelClient) InsertHotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("")
		return model.Hotel{}
	}
	log.Debug("Hotel Created: ", hotel.HotelID)
	return hotel
}

func (c *hotelClient) UpdateHotel(hotel model.Hotel) model.Hotel {
	result := Db.Save(&hotel)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return model.Hotel{}
	}
	return hotel
}

func (c *hotelClient) DeleteHotel(id string) error {
	var hotel model.Hotel
	result := Db.Delete(&hotel, "hotel_id = ?", id)
	if result.Error != nil {
		log.Debug(id)
		log.Error(result.Error.Error())
		return errors.New(result.Error.Error())
	}
	return nil
}

func (c *hotelClient) GetAvailableRooms(booking dto.CheckAvailability) float64 {
	type Result struct {
		AvailableRooms float64
	}

	var result Result

	dateIn := booking.DateIn.Format("2006-01-02")
	dateOut := booking.DateOut.Format("2006-01-02")
	hotelID := booking.HotelID

	Db.Raw(`
	SELECT (rooms - 
		(SELECT COALESCE(SUM(rooms), 0) FROM bookings
		WHERE (active = 1) AND ((date_in >= ? AND date_in < ?)
		OR (date_out > ? AND date_out <= ?)
		OR (date_in < ? AND date_out > ?))
		AND hotel_id = ?)) 
	AS available_rooms
	FROM hotels WHERE hotel_id = ? AND active = 1;
`, dateIn, dateOut, dateIn, dateOut, dateIn, dateOut, hotelID, hotelID).Scan(&result)
	log.Debug("available rooms: ", result)

	return result.AvailableRooms
}

func (c *hotelClient) GetAvailableHotels(booking dto.CheckAvailability) model.Hotels {
	var hotels model.Hotels

	dateIn := booking.DateIn.Format("2006-01-02")
	dateOut := booking.DateOut.Format("2006-01-02")
	rooms := booking.Rooms

	Db.Raw(`
	SELECT h.*, h.rooms - COALESCE((
		SELECT SUM(b.rooms) FROM bookings b
		WHERE (b.active = 1) AND (h.active = 1) 
			AND ((b.date_in >= ? AND b.date_in < ?)
			OR (b.date_out > ? AND b.date_out <= ?)
			OR (b.date_in < ? AND b.date_out > ?))
			AND h.hotel_id = b.hotel_id
	), 0) AS available_rooms
	FROM hotels h
	WHERE h.active = 1
	GROUP BY h.hotel_id
	HAVING available_rooms >= ?
`, dateIn, dateOut, dateIn, dateOut, dateIn, dateOut, rooms).Preload("Photos").Preload("Amenities").Find(&hotels)

	return hotels
}
