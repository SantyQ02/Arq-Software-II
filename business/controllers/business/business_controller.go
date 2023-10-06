package businessController

import (
	businessService "mvc-go/services/business"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"mvc-go/dto"
)

func CheckAvailability(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotelID must be a uuid"})
		return
	}
	
	checkInDate	:= c.Query("checkInDate")
	checkOutDate := c.Query("checkOutDate")

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
