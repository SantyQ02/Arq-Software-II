package app

import (
	middlewareController "mvc-go/controllers/middleware"
	"net/http"

	"mvc-go/app/urls"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	router.MaxMultipartMemory = 8 << 20

	// Users Mapping
	api := router.Group("/api")

	api.StaticFS("/public", http.Dir("static"))

	auth := api.Group("/auth")
	urls.AuthRoute(auth)

	user := api.Group("/user")
	urls.UserRoute(user)

	hotel := api.Group("/hotel")
	urls.HotelRoute(hotel)

	booking := api.Group("/booking", middlewareController.DeserializeUser())
	urls.BookingRoute(booking)

	business := api.Group("/business")
	urls.BusinessRoute(business)

	// amenitie := api.Group("/amenitie", middlewareController.DeserializeUser())
	// urls.AmenitieRoute(amenitie)

	log.Info("Finishing mappings configurations")
}
