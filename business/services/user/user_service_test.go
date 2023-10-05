/*
User Service Tests:

***GetUserById:
TestGetUserById
TestGetUserByIdErrorUserNotFound

***GetUsers:
TestGetUsers
TestGetUsersErrorDataBase

***DeleteUser:
TestDeleteUser
TestDeleteUserError500

***UpdateUser:
TestUpdateUser
TestUpdateUserErrorUserNotFound
TestUpdateUserError500

***MakeAdminUser:
TestMakeAdminUser
TestMakeAdminUserErrorUserNotFound
TestMakeAdminUserError500
*/
package userService

import (
	"errors"
	userClient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initTestClient() {
	userClient.UserClient = &userClient.UserMockClient{}
}

// GetUserById:
func TestGetUserById(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()
	now := time.Now()

	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Name Test",
		LastName:         "Last Name Test",
		Email:            "email@test.com",
		UserName:         "usertest",
		Password:         "hashedpasword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockClient.On("GetUserById", user_id.String()).Return(userModel)

	user, err := UserService.GetUserById(user_id)

	assert.Equal(t, userModel.UserID, user.UserID)
	assert.Equal(t, userModel.FirstName, user.FirstName)
	assert.Equal(t, userModel.LastName, user.LastName)
	assert.Equal(t, userModel.Email, user.Email)
	assert.Equal(t, userModel.UserName, user.UserName)
	assert.Equal(t, userModel.Role, user.Role)
	assert.Equal(t, userModel.CreatedAt, user.CreatedAt)
	assert.Equal(t, userModel.UpdatedAt, user.UpdatedAt)
	assert.Nil(t, err)
}
func TestGetUserByIdErrorUserNotFound(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()
	var time_zero time.Time

	userModel := model.User{
		UserID:           uuid.Nil,
		FirstName:        "",
		LastName:         "",
		Email:            "",
		UserName:         "",
		Password:         "",
		Role:             "",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        time_zero,
		UpdatedAt:        time_zero,
	}

	mockClient.On("GetUserById", user_id.String()).Return(userModel)

	user, err := UserService.GetUserById(user_id)

	assert.Equal(t, uuid.Nil, user.UserID)
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "User not found", err.Message())
}

// GetUsers:
func TestGetUsers(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id_1 := uuid.New()
	user_id_2 := uuid.New()
	now := time.Now()
	userModel1 := model.User{
		UserID:           user_id_1,
		FirstName:        "Name Test 1",
		LastName:         "Last Name Test 1",
		Email:            "email1@test.com",
		UserName:         "usertest1",
		Password:         "hashedpasword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	userModel2 := model.User{
		UserID:           user_id_2,
		FirstName:        "Name Test 2",
		LastName:         "Last Name Test 2",
		Email:            "email2@test.com",
		UserName:         "usertest2",
		Password:         "hashedpasword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	usersModel := model.Users{
		userModel1,
		userModel2,
	}

	mockClient.On("GetUsers").Return(usersModel)

	users, err := UserService.GetUsers()

	assert.Equal(t, 2, len(users))

	assert.Equal(t, userModel1.UserID, users[0].UserID)
	assert.Equal(t, userModel1.FirstName, users[0].FirstName)
	assert.Equal(t, userModel1.LastName, users[0].LastName)
	assert.Equal(t, userModel1.Email, users[0].Email)
	assert.Equal(t, userModel1.UserName, users[0].UserName)
	assert.Equal(t, userModel1.Role, users[0].Role)
	assert.Equal(t, userModel1.CreatedAt, users[0].CreatedAt)
	assert.Equal(t, userModel1.UpdatedAt, users[0].UpdatedAt)

	assert.Equal(t, userModel2.UserID, users[1].UserID)
	assert.Equal(t, userModel2.FirstName, users[1].FirstName)
	assert.Equal(t, userModel2.LastName, users[1].LastName)
	assert.Equal(t, userModel2.Email, users[1].Email)
	assert.Equal(t, userModel2.UserName, users[1].UserName)
	assert.Equal(t, userModel2.Role, users[1].Role)
	assert.Equal(t, userModel2.CreatedAt, users[1].CreatedAt)
	assert.Equal(t, userModel2.UpdatedAt, users[1].UpdatedAt)

	assert.Nil(t, err)
}
func TestGetUsersErrorDataBase(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)

	usersModelError := model.Users{}
	mockClient.On("GetUsers").Return(usersModelError)

	users, err := UserService.GetUsers()

	assert.Equal(t, dto.UserResponses{}, users)
	assert.Equal(t, 0, len(users))
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error getting users from database", err.Message())
}

// DeleteUser
func TestDeleteUser(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()

	mockClient.On("DeleteUser", user_id.String()).Return(nil)

	err := UserService.DeleteUser(user_id)

	assert.Nil(t, err)
}
func TestDeleteUserError500(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()

	mockClient.On("DeleteUser", user_id.String()).Return(errors.New(""))

	err := UserService.DeleteUser(user_id)

	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Something went wrong deleting user", err.Message())
}

// UpdateUser
func TestUpdateUser(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()
	now := time.Now()

	userModel := model.User{
		UserID:           user_id,
		FirstName:        "New Name",
		LastName:         "New Last Name",
		Email:            "newemail@test.com",
		UserName:         "newusertest",
		Password:         "hashedpasword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockClient.On("GetUserById", user_id.String()).Return(userModel)
	mockClient.On("UpdateUser", mock.Anything).Return(userModel)

	userDto := dto.UserResponse{
		UserID:    user_id,
		FirstName: "New Name",
		LastName:  "New Last Name",
		Email:     "newemail@test.com",
		UserName:  "newusertest",
		Role:      "user",
		UpdatedAt: now,
	}
	user, err := UserService.UpdateUser(userDto)

	assert.Nil(t, err)
	assert.Equal(t, userDto.UserID, user.UserID)
	assert.Equal(t, userDto.FirstName, user.FirstName)
	assert.Equal(t, userDto.UserName, user.UserName)
	assert.Equal(t, userDto.UpdatedAt, user.UpdatedAt)
}
func TestUpdateUserErrorUserNotFound(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()
	now := time.Now()

	mockClient.On("GetUserById", user_id.String()).Return(model.User{})

	userDto := dto.UserResponse{
		UserID:    user_id,
		FirstName: "New Name",
		LastName:  "New Last Name",
		Email:     "newemail@test.com",
		UserName:  "newusertest",
		Role:      "user",
		UpdatedAt: now,
	}
	user, err := UserService.UpdateUser(userDto)

	assert.NotNil(t, err)
	assert.Equal(t, dto.UserResponse{}, user)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "User not found", err.Message())
}
func TestUpdateUserError500(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()
	now := time.Now()

	userModel := model.User{
		UserID:           user_id,
		FirstName:        "New Name",
		LastName:         "New Last Name",
		Email:            "newemail@test.com",
		UserName:         "newusertest",
		Password:         "hashedpasword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockClient.On("GetUserById", user_id.String()).Return(userModel)
	mockClient.On("UpdateUser", mock.Anything).Return(model.User{})

	userDto := dto.UserResponse{
		UserID:    user_id,
		FirstName: "New Name",
		LastName:  "New Last Name",
		Email:     "newemail@test.com",
		UserName:  "newusertest",
		Role:      "user",
		UpdatedAt: now,
	}
	user, err := UserService.UpdateUser(userDto)

	assert.NotNil(t, err)
	assert.Equal(t, dto.UserResponse{}, user)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error updating user", err.Message())
}

// MakeAdminUser
func TestMakeAdminUser(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()
	now := time.Now()

	userModel := model.User{
		UserID:           user_id,
		FirstName:        "New Name",
		LastName:         "New Last Name",
		Email:            "newemail@test.com",
		UserName:         "newusertest",
		Password:         "hashedpasword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockClient.On("GetUserById", user_id.String()).Return(userModel)
	userModel.Role = "admin"
	mockClient.On("UpdateUser", mock.Anything).Return(userModel)

	user, err := UserService.MakeAdminUser(user_id)

	assert.Nil(t, err)
	assert.Equal(t, "admin", user.Role)
}
func TestMakeAdminUserErrorUserNotFound(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()

	mockClient.On("GetUserById", user_id.String()).Return(model.User{})

	user, err := UserService.MakeAdminUser(user_id)

	assert.NotNil(t, err)
	assert.NotEqual(t, "admin", user.Role)
	assert.Equal(t, dto.UserResponse{}, user)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "User not found", err.Message())
}
func TestMakeAdminUserError500(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	user_id := uuid.New()
	now := time.Now()

	userModel := model.User{
		UserID:           user_id,
		FirstName:        "New Name",
		LastName:         "New Last Name",
		Email:            "newemail@test.com",
		UserName:         "newusertest",
		Password:         "hashedpasword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockClient.On("GetUserById", user_id.String()).Return(userModel)
	userModel.Role = "admin"
	mockClient.On("UpdateUser", mock.Anything).Return(model.User{})

	user, err := UserService.MakeAdminUser(user_id)

	assert.NotNil(t, err)
	assert.NotEqual(t, "admin", user.Role)
	assert.Equal(t, dto.UserResponse{}, user)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error updating user", err.Message())
}
