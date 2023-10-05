package hotelController

import (
	"mvc-go/dto"
	amenitieService "mvc-go/services/amenitie"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func GetAmenitieById(c *gin.Context) {
	log.Debug("Amenitie id to load: " + c.Param("amenitieID"))

	uuid, err := uuid.Parse(c.Param("amenitieID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amenitieID must be a uuid"})
		return
	}

	amenitieDto, er := amenitieService.AmenitieService.GetAmenitieById(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amenitie": amenitieDto})
}

func GetAmenities(c *gin.Context) {
	amenitiesDto, err := amenitieService.AmenitieService.GetAmenities()
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amenities": amenitiesDto})
}

func InsertAmenitie(c *gin.Context) {

	var amenitie dto.Amenitie
	err := c.BindJSON(&amenitie)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	amenitie, er := amenitieService.AmenitieService.InsertAmenitie(amenitie)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, amenitie)
}

func UpdateAmenitie(c *gin.Context) {
	log.Debug("Amenitie id to load: " + c.Param("amenitieID"))
	uuid, err := uuid.Parse(c.Param("amenitieID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amenitieID must be a uuid"})
		return
	}

	var payload dto.Amenitie

	err = c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload.AmenitieID = uuid
	amenitie, er := amenitieService.AmenitieService.UpdateAmenitie(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amenitie": amenitie})
}

func DeleteAmenitie(c *gin.Context) {
	log.Debug("Amenitie id to load: " + c.Param("amenitieID"))

	uuid, err := uuid.Parse(c.Param("amenitieID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amenitieID must be a uuid"})
		return
	}

	er := amenitieService.AmenitieService.DeleteAmenitie(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Amenitie deleted successfully"})
}

func LoadAmenities(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("HotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	var payload dto.Amenities

	err = c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := amenitieService.AmenitieService.LoadAmenities(uuid, payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Amenities loaded successfully"})

}

func UnloadAmenities(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("HotelID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HotelID must be a uuid"})
		return
	}

	var payload dto.Amenities

	err = c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := amenitieService.AmenitieService.UnloadAmenities(uuid, payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Amenitie unloaded successfully"})

}
