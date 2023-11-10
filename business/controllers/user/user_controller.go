package userController

import (
	"mvc-go/dto"
	"mvc-go/model"
	userService "mvc-go/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func GetMe(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(model.User)
	userDto, err := userService.UserService.GetUserById(currentUser.UserID)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userDto})
}

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("userID"))

	uuid, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID must be a uuid"})
		return
	}

	userDto, er := userService.UserService.GetUserById(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userDto})
}

func GetUsers(c *gin.Context) {
	usersDto, err := userService.UserService.GetUsers()
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": usersDto})
}

func DeleteUser(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("userID"))

	uuid, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID must be a uuid"})
		return
	}

	er := userService.UserService.DeleteUser(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "User deleted successfully"})
}

func UpdateMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)

	var payload dto.UserResponse

	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if currentUser.Role != "admin" && payload.Role == "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
		return
	}
	payload.UserID = currentUser.UserID
	user, er := userService.UserService.UpdateUser(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("userID"))
	uuid, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID must be a uuid"})
		return
	}

	var payload dto.UserResponse

	err = c.BindJSON(&payload)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload.UserID = uuid
	user, er := userService.UserService.UpdateUser(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func MakeAdminUser(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("userID"))
	uuid, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userID must be a uuid"})
		return
	}

	user, er := userService.UserService.MakeAdminUser(uuid)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
