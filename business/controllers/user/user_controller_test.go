/*
***Tests for GetMe:
TestGetMe
TestGetMeErrorService

***Tests for GetUserById:
TestGetUserById
TestGetUserByIdErrorNotAdmin
TestGetUserByIdErrorInvalidUUID
TestGetUserByIdErrorService

***Tests for GetUsers:
TestGetUsers
TestGetUsersErrorNotAdmin
TestGetUsersErrorService
*/
package userController

import (
	"bytes"
	"encoding/json"
	"errors"
	middlewareController "mvc-go/controllers/middleware"
	"mvc-go/dto"
	"mvc-go/model"
	middlewareService "mvc-go/services/middleware"
	userService "mvc-go/services/user"
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
	userService.UserService = &userService.UserMockService{}
	middlewareService.MiddlewareService = &middlewareService.MiddlewareMockService{}

}

// Test for GetMe
func TestGetMe(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	user := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "user",
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

	userDto := dto.UserResponse{
		UserID:    userID,
		FirstName: "Test Name",
		LastName:  "Last Name",
		Email:     "email@test.com",
		UserName:  "usertest",
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("GetUserById", user.UserID).Return(userDto, nil)

	router := gin.Default()
	router.GET("/test/user/me", middlewareController.DeserializeUser(), GetMe)

	req, _ := http.NewRequest("GET", "/test/user/me", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	userJSON, _ := json.Marshal(userDto)
	assert.Contains(t, resp.Body.String(), string(userJSON))
}
func TestGetMeErrorService(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	user := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "user",
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

	mockUserService.On("GetUserById", user.UserID).Return(dto.UserResponse{}, e.NewNotFoundApiError("User not found"))

	router := gin.Default()
	router.GET("/test/user/me", middlewareController.DeserializeUser(), GetMe)

	req, _ := http.NewRequest("GET", "/test/user/me", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")
}

// Tests for GetUserById
func TestGetUserById(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userFoundID := uuid.New()
	userFoundDto := dto.UserResponse{
		UserID:    userFoundID,
		FirstName: "Test Name",
		LastName:  "Last Name",
		Email:     "email@test.com",
		UserName:  "usertest",
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("GetUserById", userFoundDto.UserID).Return(userFoundDto, nil)

	router := gin.Default()
	router.GET("/test/user/:userID", middlewareController.CheckAdmin(), GetUserById)

	req, _ := http.NewRequest("GET", "/test/user/"+userFoundID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	userFoundJSON, _ := json.Marshal(userFoundDto)
	assert.Contains(t, resp.Body.String(), string(userFoundJSON))
}
func TestGetUserByIdErrorNotAdmin(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userFoundID := uuid.New()

	router := gin.Default()
	router.GET("/test/user/:userID", middlewareController.CheckAdmin(), GetUserById)

	req, _ := http.NewRequest("GET", "/test/user/"+userFoundID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "Admin privileges required")
}
func TestGetUserByIdErrorInvalidUUID(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	router := gin.Default()
	router.GET("/test/user/:userID", middlewareController.CheckAdmin(), GetUserById)

	req, _ := http.NewRequest("GET", "/test/user/no-uuid-21", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "userID must be a uuid")
}
func TestGetUserByIdErrorService(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userFoundID := uuid.New()
	mockUserService.On("GetUserById", userFoundID).Return(dto.UserResponse{}, e.NewNotFoundApiError("User not found"))

	router := gin.Default()
	router.GET("/test/user/:userID", middlewareController.CheckAdmin(), GetUserById)

	req, _ := http.NewRequest("GET", "/test/user/"+userFoundID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")
}

// Tests for GetUsers
func TestGetUsers(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userID1 := uuid.New()
	userID2 := uuid.New()
	userDto1 := dto.UserResponse{
		UserID:    userID1,
		FirstName: "Test Name",
		LastName:  "Last Name",
		Email:     "email@test.com",
		UserName:  "usertest",
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}
	userDto2 := dto.UserResponse{
		UserID:    userID2,
		FirstName: "Test Name 2",
		LastName:  "Last Name 2",
		Email:     "emails@test.com",
		UserName:  "usertestaasd",
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}
	usersDto := dto.UserResponses{userDto1, userDto2}

	mockUserService.On("GetUsers").Return(usersDto, nil)

	router := gin.Default()
	router.GET("/test/user/", middlewareController.CheckAdmin(), GetUsers)

	req, _ := http.NewRequest("GET", "/test/user/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	usersJSON, _ := json.Marshal(usersDto)
	assert.Contains(t, resp.Body.String(), string(usersJSON))
}
func TestGetUsersErrorNotAdmin(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	router := gin.Default()
	router.GET("/test/user/", middlewareController.CheckAdmin(), GetUsers)

	req, _ := http.NewRequest("GET", "/test/user/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "Admin privileges required")
}
func TestGetUsersErrorService(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	mockUserService.On("GetUsers").Return(dto.UserResponses{}, e.NewInternalServerApiError("Error getting users from database", errors.New("Error in database")))

	router := gin.Default()
	router.GET("/test/user/", middlewareController.CheckAdmin(), GetUsers)

	req, _ := http.NewRequest("GET", "/test/user/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
	assert.Contains(t, resp.Body.String(), "Error getting users from database")
}

// Tests for DeleteUser
func TestDeleteUser(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userToDeleteID := uuid.New()
	mockUserService.On("DeleteUser", userToDeleteID).Return(nil)

	router := gin.Default()
	router.DELETE("/test/user/:userID", middlewareController.CheckAdmin(), DeleteUser)

	req, _ := http.NewRequest("DELETE", "/test/user/"+userToDeleteID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "User deleted successfully")
}
func TestDeleteUserErrorNotAdmin(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userToDeleteID := uuid.New()

	router := gin.Default()
	router.DELETE("/test/user/:userID", middlewareController.CheckAdmin(), DeleteUser)

	req, _ := http.NewRequest("DELETE", "/test/user/"+userToDeleteID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "Admin privileges required")
}
func TestDeleteUserErrorInvalidUUID(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	router := gin.Default()
	router.DELETE("/test/user/:userID", middlewareController.CheckAdmin(), DeleteUser)

	req, _ := http.NewRequest("DELETE", "/test/user/no-uuid-21", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "userID must be a uuid")
}
func TestDeleteUserErrorService(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userFoundID := uuid.New()
	mockUserService.On("DeleteUser", userFoundID).Return(e.NewInternalServerApiError("Something went wrong deleting user", nil))

	router := gin.Default()
	router.DELETE("/test/user/:userID", middlewareController.CheckAdmin(), DeleteUser)

	req, _ := http.NewRequest("DELETE", "/test/user/"+userFoundID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
	assert.Contains(t, resp.Body.String(), "Something went wrong deleting user")
}

// Tests for UpdateMe
func TestUpdateUserMe(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("DeserializeUser").Return(handlerFunc)

	userUpdated := dto.UserResponse{
		UserID:    userID,
		FirstName: "NuevoNombre",
		LastName:  "NuevoApellido",
		Email:     "email@test.com",
		UserName:  "nuevousername",
		Role:      "user",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("UpdateUser", mock.Anything).Return(userUpdated, nil)

	router := gin.Default()
	router.PUT("/test/user/:userID", middlewareController.DeserializeUser(), UpdateMe)

	loginPayload := []byte(`{
		"first_name": "NuevoNombre",
		"last_name": "NuevoApellido",
		"user_name": "nuevousername",
		"role": "user"
	}`)

	req, _ := http.NewRequest("PUT", "/test/user/"+userID.String(), bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	userJSON, _ := json.Marshal(userUpdated)
	assert.Contains(t, resp.Body.String(), string(userJSON))
}
func TestUpdateUserMeError(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("DeserializeUser").Return(handlerFunc)

	userUpdated := dto.UserResponse{
		UserID:    userID,
		FirstName: "NuevoNombre",
		LastName:  "NuevoApellido",
		Email:     "email@test.com",
		UserName:  "nuevousername",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("UpdateUser", mock.Anything).Return(userUpdated, nil)

	router := gin.Default()
	router.PUT("/test/user/:userID", middlewareController.DeserializeUser(), UpdateMe)

	loginPayload := []byte(`{
		"first_name": "NuevoNombre",
		"last_name": "NuevoApellido",
		"user_name": "nuevousername",
		"role": "admin"
	}`)

	req, _ := http.NewRequest("PUT", "/test/user/"+userID.String(), bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "Admin privileges required")
}

// Tests for UpdateUser
func TestUpdateUser(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	now := time.Now()

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userUpdatedID := uuid.New()
	userUpdated := dto.UserResponse{
		UserID:    userUpdatedID,
		FirstName: "NuevoNombre",
		LastName:  "NuevoApellido",
		Email:     "unemail@gmail.com",
		UserName:  "nuevousername",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("UpdateUser", mock.Anything).Return(userUpdated, nil)

	router := gin.Default()
	router.PUT("/test/user/:userID", middlewareController.CheckAdmin(), UpdateUser)

	loginPayload := []byte(`{
		"first_name": "NuevoNombre",
		"last_name": "NuevoApellido",
		"user_name": "nuevousername",
		"role": "admin"
	}`)

	req, _ := http.NewRequest("PUT", "/test/user/"+userUpdatedID.String(), bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	userJSON, _ := json.Marshal(userUpdated)
	assert.Contains(t, resp.Body.String(), string(userJSON))
}
func TestUpdateUserErrorInvalidUUID(t *testing.T) {
	initTestClient()
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)
	router := gin.Default()
	router.PUT("/test/user/:userID", middlewareController.CheckAdmin(), UpdateUser)

	req, _ := http.NewRequest("PUT", "/test/user/sfsaf124234", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "userID must be a uuid")
}
func TestUpdateUserErrorNotAdmin(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	now := time.Now()

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userUpdatedID := uuid.New()
	userUpdated := dto.UserResponse{
		UserID:    userUpdatedID,
		FirstName: "NuevoNombre",
		LastName:  "NuevoApellido",
		Email:     "unemail@gmail.com",
		UserName:  "nuevousername",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("UpdateUser", mock.Anything).Return(userUpdated, nil)

	router := gin.Default()
	router.PUT("/test/user/:userID", middlewareController.CheckAdmin(), UpdateUser)

	loginPayload := []byte(`{
		"first_name": "NuevoNombre",
		"last_name": "NuevoApellido",
		"user_name": "nuevousername",
		"role": "admin"
	}`)

	req, _ := http.NewRequest("PUT", "/test/user/"+userUpdatedID.String(), bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "Admin privileges required")
}
func TestUpdateUserErrorService(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userUpdatedID := uuid.New()

	mockUserService.On("UpdateUser", mock.Anything).Return(dto.UserResponse{}, e.NewNotFoundApiError("User not found"))

	router := gin.Default()
	router.PUT("/test/user/:userID", middlewareController.CheckAdmin(), UpdateUser)

	loginPayload := []byte(`{
		"first_name": "NuevoNombre",
		"last_name": "NuevoApellido",
		"user_name": "nuevousername",
		"role": "admin"
	}`)

	req, _ := http.NewRequest("PUT", "/test/user/"+userUpdatedID.String(), bytes.NewBuffer(loginPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")
}

// Tests for MakeAdminUser
func TestMakeAdminUser(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userUpdatedID := uuid.New()
	userUpdated := dto.UserResponse{
		UserID:    userUpdatedID,
		FirstName: "NuevoNombre",
		LastName:  "NuevoApellido",
		Email:     "unemail@gmail.com",
		UserName:  "nuevousername",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("MakeAdminUser", mock.Anything).Return(userUpdated, nil)

	router := gin.Default()
	router.PUT("/test/user/superuser/:userID", middlewareController.CheckAdmin(), MakeAdminUser)

	req, _ := http.NewRequest("PUT", "/test/user/superuser/"+userUpdatedID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	userJSON, _ := json.Marshal(userUpdated)
	assert.Contains(t, resp.Body.String(), string(userJSON))
}
func TestMakeAdminUserErrorNotAdmin(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	now := time.Now()
	userUpdatedID := uuid.New()
	userUpdated := dto.UserResponse{
		UserID:    userUpdatedID,
		FirstName: "NuevoNombre",
		LastName:  "NuevoApellido",
		Email:     "unemail@gmail.com",
		UserName:  "nuevousername",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("MakeAdminUser", mock.Anything).Return(userUpdated, nil)

	router := gin.Default()
	router.PUT("/test/user/superuser/:userID", middlewareController.CheckAdmin(), MakeAdminUser)

	req, _ := http.NewRequest("PUT", "/test/user/superuser/"+userUpdatedID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 401, resp.Code)
	assert.Contains(t, resp.Body.String(), "Admin privileges required")
}
func TestMakeAdminUserErrorInvalidUUID(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userUpdatedID := uuid.New()
	userUpdated := dto.UserResponse{
		UserID:    userUpdatedID,
		FirstName: "NuevoNombre",
		LastName:  "NuevoApellido",
		Email:     "unemail@gmail.com",
		UserName:  "nuevousername",
		Role:      "admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
	mockUserService.On("MakeAdminUser", mock.Anything).Return(userUpdated, nil)

	router := gin.Default()
	router.PUT("/test/user/superuser/:userID", middlewareController.CheckAdmin(), MakeAdminUser)

	req, _ := http.NewRequest("PUT", "/test/user/superuser/sdfas1234fsadf", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
	assert.Contains(t, resp.Body.String(), "userID must be a uuid")
}
func TestMakeAdminUserErrorService(t *testing.T) {
	initTestClient()
	mockUserService := userService.UserService.(*userService.UserMockService)
	mockMiddleware := middlewareService.MiddlewareService.(*middlewareService.MiddlewareMockService)

	userID := uuid.New()
	now := time.Now()
	userAdmin := model.User{
		UserID:           userID,
		FirstName:        "Test Name",
		LastName:         "Last Name",
		Email:            "email@test.com",
		UserName:         "usertest",
		Role:             "admin",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	var handlerFunc gin.HandlerFunc
	handlerFunc = func(c *gin.Context) {
		c.Set("currentUser", userAdmin)
		c.Next()
	}
	mockMiddleware.On("CheckAdmin").Return(handlerFunc)

	userUpdatedID := uuid.New()

	mockUserService.On("MakeAdminUser", mock.Anything).Return(dto.UserResponse{}, e.NewInternalServerApiError("Error updating user", nil))

	router := gin.Default()
	router.PUT("/test/user/superuser/:userID", middlewareController.CheckAdmin(), MakeAdminUser)

	req, _ := http.NewRequest("PUT", "/test/user/superuser/"+userUpdatedID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
	assert.Contains(t, resp.Body.String(), "Error updating user")
}
