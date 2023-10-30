package authController

import (
	"mvc-go/dto"
	"mvc-go/model"
	authService "mvc-go/services/auth"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

// [...] SignUp User
func RegisterUser(c *gin.Context) {
	var payload dto.Register

	err := c.BindJSON(&payload)
	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerDto, er := authService.AuthService.RegisterUser(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	message := "We sent an email with a verification code to " + registerDto.Email
	c.JSON(http.StatusCreated, gin.H{"success": message})
}

// [...] Verify Email
func VerifyEmail(c *gin.Context) {

	code := c.Params.ByName("verificationCode")

	err := authService.AuthService.VerifyEmail(code)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Email verified successfully"})
}

// [...] SignIn User
func LoginUser(c *gin.Context) {
	var payload dto.Login

	err := c.BindJSON(&payload)
	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// var errorEmpty e.ApiError
	token, er := authService.AuthService.LoginUser(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	tokenMaxAge, _ := strconv.Atoi(os.Getenv("TOKEN_MAXAGE"))
	c.SetCookie("token", token, tokenMaxAge*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Refresh(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)
	// var errorEmpty e.ApiError
	token, er := authService.AuthService.Refresh(currentUser.UserID)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	tokenMaxAge, _ := strconv.Atoi(os.Getenv("TOKEN_MAXAGE"))
	c.SetCookie("token", token, tokenMaxAge*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// [...] SignOut User
func LogoutUser(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"success": "Logged out successfully"})
}

func ResetPassword(c *gin.Context) {
	var payload dto.ResetPassword

	err := c.BindJSON(&payload)
	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resetPassDto, er := authService.AuthService.ResetPassword(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	message := "We sent an email to reset your password " + resetPassDto.Email
	c.JSON(http.StatusOK, gin.H{"success": message})
}

func ResetPasswordConfirm(c *gin.Context) {
	code := c.Params.ByName("verificationCode")
	var payload dto.ResetPasswordConfirm

	err := c.BindJSON(&payload)
	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload.VerificationCode = code

	er := authService.AuthService.ResetPasswordConfirm(payload)
	if er != nil {
		c.JSON(er.Status(), gin.H{"error": er.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Password reset successfully"})
}
