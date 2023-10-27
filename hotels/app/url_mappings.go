package app

import (
	"net/http"

	"mvc-go/app/urls"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	router.MaxMultipartMemory = 8 << 20

	// Users Mapping
	api := router.Group("/api")

	api.StaticFS("/public", http.Dir("static"))

	hotel := api.Group("/hotel")
	urls.HotelRoute(hotel)


	// amenitie := api.Group("/amenitie", middlewareController.DeserializeUser())
	// urls.AmenitieRoute(amenitie)

	log.Info("Finishing mappings configurations")
}
