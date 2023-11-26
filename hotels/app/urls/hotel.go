package urls

import (
	hotelController "mvc-go/controllers/hotel"
	middlewareController "mvc-go/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func HotelRoute(hotel *gin.RouterGroup) {
	
	hotel.GET("/:hotelID", hotelController.GetHotelById)
	hotel.GET("/", hotelController.GetHotels)

	// Only admin:
	hotel.POST("/",middlewareController.CheckAdmin(),hotelController.InsertHotel)
	hotel.PUT("/:hotelID",middlewareController.CheckAdmin(), hotelController.UpdateHotel)
	hotel.DELETE("/:hotelID",middlewareController.CheckAdmin(), hotelController.DeleteHotel)
	hotel.POST("/upload/:hotelID",middlewareController.CheckAdmin(), hotelController.UploadPhoto)

}
