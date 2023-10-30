package urls

import (
	authController "mvc-go/controllers/auth"
	middlewareController "mvc-go/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoute(auth *gin.RouterGroup) {
	auth.POST("/register", authController.RegisterUser)
	auth.POST("/login", authController.LoginUser)
	auth.GET("/refresh", middlewareController.DeserializeUser(), authController.Refresh)
	auth.GET("/logout", middlewareController.DeserializeUser(), authController.LogoutUser)
	auth.GET("/verifyemail/:verificationCode", authController.VerifyEmail)
	auth.POST("/reset-password", authController.ResetPassword)
	auth.POST("/reset-password/:verificationCode", authController.ResetPasswordConfirm)

}
