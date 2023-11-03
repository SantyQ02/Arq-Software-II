package bookingController

import (
	"mvc-go/dto"

	"time"

	"mvc-go/model"

	bookingService "mvc-go/services/booking"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	//"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func CreateBooking(c *gin.Context) {
	var payload dto.Booking
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	booking, er := bookingService.BookingService.CreateBooking(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"booking": booking})
}

func GetBookingsMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)

	bookingsDto, er := bookingService.BookingService.GetBookingsByUserId(currentUser.UserID)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookingsDto})

}

func SearchBookingsMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)
	var err error

	date_in := c.Query("date_in")
	var dateIn time.Time
	if date_in == "" {
		dateIn, err = time.Parse("2006-01-02T15:04:05.000Z", "1900-01-01T00:00:00.000Z")
	} else {
		dateIn, err = time.Parse("2006-01-02T15:04:05.000Z", date_in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "date_in must be a correct date"})
			return
		}
	}

	date_out := c.Query("date_out")
	var dateOut time.Time
	if date_in == "" {
		dateOut, err = time.Parse("2006-01-02T15:04:05.000Z", "2900-01-01T00:00:00.000Z")
	} else {
		dateOut, err = time.Parse("2006-01-02T15:04:05.000Z", date_out)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "date_out must be a correct date"})
			return
		}
	}

	hotel := c.Query("hotel")
	user := currentUser.UserID.String()
	log.Debug("user: ", user)
	bookings, er := bookingService.BookingService.SearchBookings(hotel, user, dateIn, dateOut)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func GetBookings(c *gin.Context) {

	bookingsDto, err := bookingService.BookingService.GetBookings()
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookingsDto})
}

func SearchBookings(c *gin.Context) {
	var err error

	date_in := c.Query("date_in")
	var dateIn time.Time
	if date_in == "" {
		dateIn, err = time.Parse("2006-01-02T15:04:05.000Z", "1900-01-01T00:00:00.000Z")
	} else {
		dateIn, err = time.Parse("2006-01-02T15:04:05.000Z", date_in)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "date_in must be a correct date"})
			return
		}
	}

	date_out := c.Query("date_out")
	var dateOut time.Time
	if date_in == "" {
		dateOut, err = time.Parse("2006-01-02T15:04:05.000Z", "2900-01-01T00:00:00.000Z")
	} else {
		dateOut, err = time.Parse("2006-01-02T15:04:05.000Z", date_out)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "date_out must be a correct date"})
			return
		}
	}

	hotel := c.Query("hotel")
	user := c.Query("user")
	bookings, er := bookingService.BookingService.SearchBookings(hotel, user, dateIn, dateOut)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

func GetBookingById(c *gin.Context) {
	
}

func SetActiveBooking(c *gin.Context) {
	log.Debug("Booking id to set active: " + c.Param("bookingID"))
	uuid, err := uuid.Parse(c.Param("bookingID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookingID must be an uuid"})
		return
	}

	booking, er := bookingService.BookingService.GetBookingById(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	if currentUser.Role != "admin" && currentUser.UserID != booking.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permissions to delete this booking"})
		return
	}

	var payload dto.SetActive
	err = c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	booking.Active = payload.Active

	booking, er = bookingService.BookingService.SetActiveBooking(booking)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})

}

func DeleteBooking(c *gin.Context) {
	log.Debug("Booking id to delete: " + c.Param("bookingID"))
	uuid, err := uuid.Parse(c.Param("bookingID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookingID must be an uuid"})
		return
	}

	booking, er := bookingService.BookingService.GetBookingById(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	if currentUser.Role != "admin" && currentUser.UserID != booking.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permissions to delete this booking"})
		return
	}

	er = bookingService.BookingService.DeleteBooking(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Booking deleted successfully"})

}
