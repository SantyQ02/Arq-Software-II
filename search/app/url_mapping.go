package app

import (
	"mvc-go/controller"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	router.MaxMultipartMemory = 8 << 20

	api := router.Group("/api")

	api.GET("/search", controller.Search)

	log.Info("Finishing mappings configurations")
}