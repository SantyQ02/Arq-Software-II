package businessController

import (
	"mvc-go/dto"
	businessService "mvc-go/services/business"
	"net/http"
	"os"
	"time"

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

func CheckAdmin(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID must be a uuid"})
		return
	}

	queryKey := c.Query("key")
	key := os.Getenv("QUERY_KEY")
	if queryKey != key {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid key"})
		return
	}

	adminbool, er := businessService.BusinessService.CheckAdmin(uuid)
	if er != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": uuid, "admin": adminbool})
}
