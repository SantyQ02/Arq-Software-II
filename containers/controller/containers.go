package controller

import (
	"net/http"
	"containers/dto"
	"containers/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


func GetContainersStats(c *gin.Context) {
	containers_stats, er := service.ContainersService.GetContainersStats()
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, containers_stats)
}

func CreateContainer(c *gin.Context) {
	var payload dto.CreateContainer
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	er := service.ContainersService.CreateContainer(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Container created successfully"})
}

func DeleteContainer(c *gin.Context) {

	container_id := c.Param("container_id")

	er := service.ContainersService.DeleteContainer(container_id)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Container deleted successfully"})

}

func RestartContainer(c *gin.Context) {

	container_id := c.Param("container_id")

	er := service.ContainersService.RestartContainer(container_id)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Container restarted successfully"})

}
