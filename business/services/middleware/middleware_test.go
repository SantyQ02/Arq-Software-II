/*
 */
package middlewareService

import (
	"encoding/json"
	"errors"
	userClient "mvc-go/clients/user"
	"mvc-go/model"
	authUtils "mvc-go/utils/auth"
	"mvc-go/utils/email"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initTestClient() {
	// AuthService = &authService{}

	userClient.UserClient = &userClient.UserMockClient{}
	email.EmailClient = &email.EmailMockClient{}
	authUtils.PassClient = &authUtils.PassMockClient{}
	authUtils.TokenClient = &authUtils.TokenMockClient{}

}

func TestDeserializeUserWithLocalStorage(t *testing.T) {
	initTestClient()
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            "email@test.com",
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockTokenClient.On("ValidateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("a-uuid", nil)

	mockUserClient := userClient.UserClient.(*userClient.UserMockClient)
	mockUserClient.On("GetUserById", "a-uuid").Return(userModel)

	router := gin.Default()
	router.GET("/test/middleware", MiddlewareService.DeserializeUser(), func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"user": c.MustGet("currentUser").(model.User)}) })

	req, _ := http.NewRequest("GET", "/test/middleware", nil)
	req.Header.Set("Authorization", "Bearer your-token-here")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	userJSON, _ := json.Marshal(userModel)
	assert.Contains(t, resp.Body.String(), string(userJSON))
}
func TestDeserializeUserWithCookie(t *testing.T) {
	initTestClient()
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            "email@test.com",
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockTokenClient.On("ValidateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("a-uuid", nil)

	mockUserClient := userClient.UserClient.(*userClient.UserMockClient)
	mockUserClient.On("GetUserById", "a-uuid").Return(userModel)

	router := gin.Default()
	router.GET("/test/middleware", MiddlewareService.DeserializeUser(), func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"user": c.MustGet("currentUser").(model.User)}) })

	req, _ := http.NewRequest("GET", "/test/middleware", nil)
	cookie := &http.Cookie{
		Name:  "token",
		Value: "a-jwt",
	}
	req.AddCookie(cookie)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	userJSON, _ := json.Marshal(userModel)
	assert.Contains(t, resp.Body.String(), string(userJSON))
}
func TestDeserializeUserErrorNotLoggedIn(t *testing.T) {
	initTestClient()
	router := gin.Default()
	router.GET("/test/middleware", MiddlewareService.DeserializeUser(), func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Deserialized Succesfully"}) })

	req, _ := http.NewRequest("GET", "/test/middleware", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "You are not logged in")
}
func TestDeserializeUserErrorInvalidToken(t *testing.T) {
	initTestClient()
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)
	mockTokenClient.On("ValidateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errors.New(""))

	router := gin.Default()
	router.GET("/test/middleware", MiddlewareService.DeserializeUser(), func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Deserialized Succesfully"}) })

	req, _ := http.NewRequest("GET", "/test/middleware", nil)
	cookie := &http.Cookie{
		Name:  "token",
		Value: "a-jwt",
	}
	req.AddCookie(cookie)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
}
func TestDeserializeUserErrorForbidden(t *testing.T) {
	initTestClient()
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)

	mockTokenClient.On("ValidateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("a-uuid", nil)

	mockUserClient := userClient.UserClient.(*userClient.UserMockClient)
	mockUserClient.On("GetUserById", "a-uuid").Return(model.User{})

	router := gin.Default()
	router.GET("/test/middleware", MiddlewareService.DeserializeUser(), func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Deserialized Succesfully"}) })

	req, _ := http.NewRequest("GET", "/test/middleware", nil)
	req.Header.Set("Authorization", "Bearer your-token-here")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 403, resp.Code)
	assert.Contains(t, resp.Body.String(), "the user belonging to this token no logger exists")
}
