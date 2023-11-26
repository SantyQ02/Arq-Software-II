package app

import (
	"containers/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	containers := router.Group("/api/containers")

	containers.GET("/", controller.GetContainersStats)
	containers.POST("/", controller.CreateContainer)
	
	containers.DELETE("/delete/:container_id", controller.DeleteContainer)
	containers.GET("/start/:container_id", controller.StartContainer)
	containers.GET("/stop/:container_id", controller.StopContainer)
	containers.GET("/restart/:container_id", controller.RestartContainer)

	log.Info("Finishing mappings configurations")
}