package controller

import (
	"net/http"
	// "strconv"
	"mvc-go/service"

	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"
)

func Search(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city must be a valid value"})
		return
	}
	checkInDate := c.Query("start_date")
	if checkInDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start date must be a valid value"})
		return
	}
	checkOutDate := c.Query("end_date")
	if checkOutDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end date must be a valid value"})
		return
	}
	hotels, err := service.SearchService.Search(city, checkInDate, checkOutDate)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotels": hotels})
}