package hotelController

import (
	"mvc-go/dto"
	photoService "mvc-go/services/photo"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func GetPhotosByHotelId(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("hotelID"))

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotelID must be a uuid"})
		return
	}

	photosDto, er := photoService.PhotoService.GetPhotosByHotelId(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photosDto})
}

func InsertPhoto(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("hotelID"))

	uuid, err := uuid.Parse(c.Param("hotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotelID must be a uuid"})
		return
	}

	var payload dto.Photo
	err = c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload.HotelID = uuid
	photo, er := photoService.PhotoService.InsertPhoto(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"photo": photo})
}

func ChangePhoto(c *gin.Context) {

}

func DeletePhoto(c *gin.Context) {
	log.Debug("Photo id to load: " + c.Param("photoID"))

	uuid, err := uuid.Parse(c.Param("photoID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "photoID must be a uuid"})
		return
	}

	er := photoService.PhotoService.DeletePhoto(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Photo deleted successfully"})

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
	payload.HotelID = uuid

	photo, er := photoService.PhotoService.InsertPhoto(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"photo": photo})
}
