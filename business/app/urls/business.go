package urls

import (
	businessController "mvc-go/controllers/business"
	middlewareController "mvc-go/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func BusinessRoute(business *gin.RouterGroup) {
	business.GET("/availability/:hotelID", businessController.CheckAvailability)
	business.POST("/mapping-hotel", businessController.MapHotel)
	business.GET("/checkadmin", middlewareController.CheckAdmin(), businessController.CheckAdmin)
}
