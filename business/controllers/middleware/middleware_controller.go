package middlewareController

import (
	middlewareService "mvc-go/services/middleware"

	"github.com/gin-gonic/gin"
)

func DeserializeUser() gin.HandlerFunc {
	return middlewareService.MiddlewareService.DeserializeUser()
}

func CheckAdmin() gin.HandlerFunc {
	return middlewareService.MiddlewareService.CheckAdmin()
}
