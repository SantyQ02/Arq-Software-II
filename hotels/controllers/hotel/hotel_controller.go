package hotelController

import (
	// "mvc-go/model"
	// userService "mvc-go/services/user"
	"net/http"
	"path/filepath"

	"mvc-go/dto"
	hotelService "mvc-go/services/hotel"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


func GetHotelById(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	hotel, er := hotelService.HotelService.GetHotelById(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotel": hotel})

}

func InsertHotel(c *gin.Context) {
	var payload dto.Hotel
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotel, er := hotelService.HotelService.InsertHotel(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"hotel": hotel})
}

func UpdateHotel(c *gin.Context) {

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	var payload dto.Hotel
	errr := c.BindJSON(&payload)
	if errr != nil {
		c.JSON(http.StatusBadRequest, errr.Error())
		return
	}

	payload.HotelID = uuid

	hotel, er := hotelService.HotelService.UpdateHotel(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hotel": hotel})
}

func DeleteHotel(c *gin.Context) {

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	er := hotelService.HotelService.DeleteHotel(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Hotel deleted successfully"})

}

func UploadPhoto(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("hotelID"))
	newFileName := uuid.New().String()

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotelID must be a uuid"})
		return
	}

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileExt := filepath.Ext(file.Filename)

	file.Filename = "/images/hotels/" + newFileName + fileExt
	log.Debug("file name: ", file.Filename)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "static"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var payload dto.Photo
	payload.Url = file.Filename

	photo, er := hotelService.HotelService.InsertPhoto(payload, uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"photo": photo})
}
