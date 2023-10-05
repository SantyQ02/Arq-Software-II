package urls

import (
	middlewareController "mvc-go/controllers/middleware"
	userController "mvc-go/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoute(user *gin.RouterGroup) {
	user.GET("/me", middlewareController.DeserializeUser(), userController.GetMe)
	user.PUT("/me", middlewareController.DeserializeUser(), userController.UpdateMe)

	user.PUT("/:userID", middlewareController.CheckAdmin(), userController.UpdateUser)
	user.GET("/:userID", middlewareController.CheckAdmin(), userController.GetUserById)
	user.GET("/", middlewareController.CheckAdmin(), userController.GetUsers)
	user.DELETE("/:userID", middlewareController.CheckAdmin(), userController.DeleteUser)
	user.PUT("/superuser/:userID", middlewareController.CheckAdmin(), userController.MakeAdminUser)

}
