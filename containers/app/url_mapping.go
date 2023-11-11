package app

import (
	"containers/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	containers := router.Group("/api/containers")

	containers.GET("/", controller.GetContainersStats)
	containers.POST("/", controller.CreateContainer)
	containers.DELETE("/:container_id", controller.DeleteContainer)
	containers.GET("/restart/:container_id", controller.RestartContainer)

	log.Info("Finishing mappings configurations")
}