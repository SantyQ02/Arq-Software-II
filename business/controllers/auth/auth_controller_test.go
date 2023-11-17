/*

***Tests for RegisterUser:
TestRegisterUser
TestRegisterUserError400
TestRegisterUserErrorService

***Tests for LoginUser:
TestLoginUser
TestLoginUserError400
TestLoginUserErrorService

***Test for LogoutUser:
TestLogout

***Tests for VerifyEmail:
TestVerifyEmail
TestVerifyEmailErrorService

***Tests for Resetpassword:
TestResetPassword
TestResetPasswordError400
TestResetPasswordErrorService

***Test for ResetPasswordConfirm:
TestResetPasswordConfirm
TestResetPasswordConfirmError400
TestResetPasswordConfirmErrorService

*/

package authController

import (
	"bytes"
	"errors"
	middlewareController "mvc-go/controllers/middleware"
	"mvc-go/dto"
	"mvc-go/model"
	authService "mvc-go/services/auth"
	middlewareService "mvc-go/services/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	e "mvc-go/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initTestClient() {
	authService.AuthService = &authService.AuthMockService{}
	middlewareService.MiddlewareService = &middlewareService.MiddlewareMockService{}

}

// Tests for Registeruser
func TestRegisterUser(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	registerDto := dto.Register{
		FirstName:       "Giovanni",
		LastName:        "Borgogno",
		Email:           "test@gmail.com",
		UserName:        "eri",
		Password:        "pAss12345",
		PasswordConfirm: "pAss12345",
	}
	mockAuthService.On("RegisterUser", registerDto).Return(registerDto, nil)

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/register", RegisterUser)
	// Crea un registro de prueba
	loginPayload := []byte(`{
		"first_name": "Giovanni",
		"last_name": "Borgogno",
		"email": "test@gmail.com",
		"user_name": "eri",
		"password": "pAss12345",
		"password_confirm" : "pAss12345"
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/register", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 201, resp.Code)
	assert.Contains(t, resp.Body.String(), "We sent an email with a verification code to "+registerDto.Email)
}
func TestRegisterUserError400(t *testing.T) {
	{
		initTestClient()

		// Crea un router Gin para la prueba
		router := gin.Default()
		router.POST("/test/register", RegisterUser)
		// Crea un registro de prueba
		loginPayload := []byte(`{
		"first_name": "Giovanni",
		"last_name": "Borgogno",
		"email": "test@gmail.com",
		"user_name": "eri",
		"password": "pAss12345"
	}`)

		// Realiza una solicitud HTTP de prueba con el registro de prueba
		req, _ := http.NewRequest("POST", "/test/register", bytes.NewBuffer(loginPayload))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, 400, resp.Code)
	}
}
func TestRegisterUserErrorService(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	registerDto := dto.Register{
		FirstName:       "Giovanni",
		LastName:        "Borgogno",
		Email:           "test@gmail.com",
		UserName:        "eri",
		Password:        "pAss12345",
		PasswordConfirm: "pAss12345",
	}
	mockAuthService.On("RegisterUser", registerDto).Return(dto.Register{}, e.NewInternalServerApiError("Error hashing password", errors.New("")))

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/register", RegisterUser)
	// Crea un registro de prueba
	loginPayload := []byte(`{
		"first_name": "Giovanni",
		"last_name": "Borgogno",
		"email": "test@gmail.com",
		"user_name": "eri",
		"password": "pAss12345",
		"password_confirm" : "pAss12345"
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/register", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
	assert.Contains(t, resp.Body.String(), "Error hashing password")
}

// Tests for LoginUser
func TestLoginUserWithEmail(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	mockAuthService.On("LoginUser", mock.Anything).Return("JWTTest", nil)

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/login", LoginUser)
	// Crea un registro de prueba
	loginPayload := []byte(`{
			"email": "tEst@example.com",
			"password": "password123"
		}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/login", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "JWTTest")
}
func TestLoginUserWithUserName(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	mockAuthService.On("LoginUser", mock.Anything).Return("JWTTest", nil)

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/login", LoginUser)
	// Crea un registro de prueba
	loginPayload := []byte(`{
			"user_name": "UserName",
			"password": "password123"
		}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/login", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "JWTTest")
}
func TestLoginUserError400(t *testing.T) {
	initTestClient()
	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/login", LoginUser)
	// Crea un registro de prueba
	loginPayload := []byte(`{
			"email": "test@example.com"
		}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/login", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}
func TestLoginUserErrorService(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	mockAuthService.On("LoginUser", mock.Anything).Return("", e.NewUnauthorizedApiError("Invalid email or Password"))

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/login", LoginUser)
	// Crea un registro de prueba
	loginPayload := []byte(`{
			"email": "test@example.com",
			"password": "password123"
		}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/login", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid email or Password")
}

// Tests for Refresh
func TestResfresh(t *testing.T) {
	initTestClient()

	userID := uuid.New()
	now := time.Now()
	user := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", user)
		c.Next()
	}
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)
	mockMiddleware.On("DeserializeUser").Return(handlerFunc)

	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	mockAuthService.On("Refresh", mock.Anything).Return("JWTTest", nil)

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.GET("/test/refresh", middlewareController.DeserializeUser(), Refresh)

	req, _ := http.NewRequest("GET", "/test/refresh", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "JWTTest")
}

// Test for LogoutUser
func TestLogout(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	user := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", user)
		c.Next()
	}

	mockMiddleware.On("DeserializeUser").Return(handlerFunc)

	router := gin.Default()
	router.GET("/test/logout", middlewareController.DeserializeUser(), LogoutUser)

	req, _ := http.NewRequest("GET", "/test/logout", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Logged out successfully")
}

// Tests for VerifyEmail
func TestVerifyEmail(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	mockAuthService.On("VerifyEmail", "sdfs1d56sd57").Return(nil)

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.GET("/test/verifyemail/:verificationCode", VerifyEmail)
	// Crea un registro de prueba

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("GET", "/test/verifyemail/sdfs1d56sd57", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Email verified successfully")
}
func TestVerifyEmailErrorService(t *testing.T) {
	{
		initTestClient()
		mockAuthService := authService.AuthService.(*authService.AuthMockService)
		mockAuthService.On("VerifyEmail", "sdfs1d56sd57").Return(e.NewNotFoundApiError("Invalid verification code or user doesn't exists"))

		// Crea un router Gin para la prueba
		router := gin.Default()
		router.GET("/test/verifyemail/:verificationCode", VerifyEmail)
		// Crea un registro de prueba

		// Realiza una solicitud HTTP de prueba con el registro de prueba
		req, _ := http.NewRequest("GET", "/test/verifyemail/sdfs1d56sd57", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, 404, resp.Code)
		assert.Contains(t, resp.Body.String(), "Invalid verification code or user doesn't exists")
	}
}

// Tests for ResetPassword
func TestResetPassword(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	resetDto := dto.ResetPassword{
		Email: "test@gmail.com",
	}
	mockAuthService.On("ResetPassword", resetDto).Return(resetDto, nil)

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/reset-password", ResetPassword)
	// Crea un registro de prueba
	loginPayload := []byte(`{
		"email": "test@gmail.com"
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/reset-password", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "We sent an email to reset your password "+resetDto.Email)
}
func TestResetPasswordError400(t *testing.T) {
	initTestClient()
	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/reset-password", ResetPassword)
	// Crea un registro de prueba
	loginPayload := []byte(`{
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/reset-password", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}
func TestResetPasswordErrorService(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	resetDto := dto.ResetPassword{
		Email: "test@gmail.com",
	}
	mockAuthService.On("ResetPassword", resetDto).Return(dto.ResetPassword{}, e.NewNotFoundApiError("Invalid email"))

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/reset-password", ResetPassword)
	// Crea un registro de prueba
	loginPayload := []byte(`{
		"email": "test@gmail.com"
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/reset-password", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid email")
}

// Tests for ResetPasswordConfirm
func TestResetPasswordConfirm(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	resetDto := dto.ResetPasswordConfirm{
		Password:         "pAss12345",
		PasswordConfirm:  "pAss12345",
		VerificationCode: "asdfas7d6fy",
	}
	mockAuthService.On("ResetPasswordConfirm", resetDto).Return(nil)

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/reset-password/:verificationCode", ResetPasswordConfirm)
	// Crea un registro de prueba
	loginPayload := []byte(`{
		"password": "pAss12345",
		"password_confirm" : "pAss12345"
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/reset-password/asdfas7d6fy", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "Password reset successfully")
}
func TestResetPasswordConfirmError400(t *testing.T) {
	initTestClient()
	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/reset-password/:verificationCode", ResetPasswordConfirm)
	// Crea un registro de prueba
	loginPayload := []byte(`{
		"password": "pAss12345"
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/reset-password/asdfas7d6fy", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}
func TestResetPasswordConfirmErrorService(t *testing.T) {
	initTestClient()
	mockAuthService := authService.AuthService.(*authService.AuthMockService)
	resetDto := dto.ResetPasswordConfirm{
		Password:         "pAss12345",
		PasswordConfirm:  "pAss12345",
		VerificationCode: "asdfas7d6fy",
	}
	mockAuthService.On("ResetPasswordConfirm", resetDto).Return(e.NewNotFoundApiError("Invalid verification code"))

	// Crea un router Gin para la prueba
	router := gin.Default()
	router.POST("/test/reset-password/:verificationCode", ResetPasswordConfirm)
	// Crea un registro de prueba
	loginPayload := []byte(`{
		"password": "pAss12345",
		"password_confirm" : "pAss12345"
	}`)

	// Realiza una solicitud HTTP de prueba con el registro de prueba
	req, _ := http.NewRequest("POST", "/test/reset-password/asdfas7d6fy", bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid verification code")
}
