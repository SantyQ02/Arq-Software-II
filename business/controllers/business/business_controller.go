package businessController

import (
	businessService "mvc-go/services/business"
	"net/http"
	"time"

	"mvc-go/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func CheckAvailability(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotelID must be a uuid"})
		return
	}

	checkInDate, erIn := time.Parse("2006-01-02", c.Query("checkInDate"))
	if erIn != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkInDate must be a valid value"})
		return
	}
	checkOutDate, erOut := time.Parse("2006-01-02", c.Query("checkOutDate"))
	if erOut != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checkOutDate must be a valid value"})
		return
	}

	availability, er := businessService.BusinessService.CheckAvailability(uuid, checkInDate, checkOutDate)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotel_id": uuid, "available": availability})
}

func MapHotel(c *gin.Context) {
	var payload dto.HotelMapping
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotel_mapping, er := businessService.BusinessService.MapHotel(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"hotel_mapping": hotel_mapping})
}
