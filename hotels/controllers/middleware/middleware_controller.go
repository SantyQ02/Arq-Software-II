package middlewareController

import (
	middlewareService "mvc-go/services/middleware_services"

	"github.com/gin-gonic/gin"
)

func CheckAdmin() gin.HandlerFunc {
	return middlewareService.MiddlewareService.CheckAdmin()
}