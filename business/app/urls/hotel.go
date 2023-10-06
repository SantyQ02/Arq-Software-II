package urls

import (
	hotelController "mvc-go/controllers/hotel"
	middlewareController "mvc-go/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func HotelRoute(hotel *gin.RouterGroup) {
	hotel.GET("/", hotelController.GetHotels)
	hotel.GET("/:hotelID", hotelController.GetHotelById)
	hotel.GET("/available", hotelController.GetAvailableHotels)
	hotel.GET("/availableRooms", hotelController.GetAvailableRooms)

	// Only admin:
	hotel.POST("/", middlewareController.CheckAdmin(), hotelController.InsertHotel)
	hotel.PUT("/:hotelID", middlewareController.CheckAdmin(), hotelController.UpdateHotel)
	hotel.DELETE("/:hotelID", middlewareController.CheckAdmin(), hotelController.DeleteHotel)

	photo := hotel.Group("/photo")
	photo.GET("/hotel/:hotelID", hotelController.GetPhotosByHotelId)

	// Only admin:
	photo.POST("/:hotelID", middlewareController.CheckAdmin(), hotelController.InsertPhoto)
	photo.DELETE("/:photoID", middlewareController.CheckAdmin(), hotelController.DeletePhoto)
	photo.POST("/upload/:hotelID", middlewareController.CheckAdmin(), hotelController.UploadPhoto)

	amenitie := hotel.Group("/amenitie")
	amenitie.GET("/", hotelController.GetAmenities)

	// Only admin:
	amenitie.POST("/", middlewareController.CheckAdmin(), hotelController.InsertAmenitie)
	amenitie.PUT("/:amenitieID", middlewareController.CheckAdmin(), hotelController.UpdateAmenitie)
	amenitie.DELETE("/:amenitieID", middlewareController.CheckAdmin(), hotelController.DeleteAmenitie)
	amenitie.POST("/loadamenities/:HotelID", middlewareController.CheckAdmin(), hotelController.LoadAmenities)
	amenitie.PUT("/unloadamenities/:HotelID", middlewareController.CheckAdmin(), hotelController.UnloadAmenities)
}