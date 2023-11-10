package middlewareController

import (
	"mvc-go/model"
	middlewareService "mvc-go/services/middleware"
	"testing"

	"github.com/gin-gonic/gin"
)

func initTestClient() {
	middlewareService.MiddlewareService = &middlewareService.MiddlewareMockService{}

}

func TestDeserializeUser(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", model.User{})
		c.Next()
	}

	mockMiddleware.On("DeserializeUser").Return(handlerFunc)
}
