package urls

import (
	hotelController "mvc-go/controllers/hotel"

	"github.com/gin-gonic/gin"
)

func HotelRoute(hotel *gin.RouterGroup) {
	
	hotel.GET("/:hotelID", hotelController.GetHotelById)

	// Only admin:
	hotel.POST("/", hotelController.InsertHotel)
	hotel.PUT("/:hotelID", hotelController.UpdateHotel)
	hotel.DELETE("/:hotelID", hotelController.DeleteHotel)
	hotel.POST("/upload/:hotelID", hotelController.UploadPhoto)

}
