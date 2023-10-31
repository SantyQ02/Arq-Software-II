package controller

import (
	"net/http"
	// "strconv"
	"mvc-go/service"

	"github.com/gin-gonic/gin"
	"time"
	// log "github.com/sirupsen/logrus"
)

func Search(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city must be a valid value"})
		return
	}
	checkInDate, err := time.Parse("2006-01-02", c.Query("check_in_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check_in_date must be a valid value"})
		return
	}
	checkOutDate, err := time.Parse("2006-01-02", c.Query("check_out_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check_out_date must be a valid value"})
		return
	}
	if checkInDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You should not have a checkInDate earlier than the current date"})
	}

	if checkInDate.After(checkOutDate) || checkInDate.Equal(checkOutDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You should not have a checkInDate greater or equal than the checkOutDate"})
	}

	hotels, er := service.SearchService.Search(city, checkInDate, checkOutDate)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotels": hotels})
}