/*
Auth Service Tests:

***RegisterUser:
TestRegisterUser
TestRegisterUserErrorPasswordsDontMatch
TestRegisterUserErrorHashingPassword
TestRegisterUserErrorUserAlreadyExists

***LoginUser:
TestLoginUser
TestLoginUserErrorNotVerified
TestLoginUserErrorInvalidUser
TestLoginUserErrorInvalidPassword
TestLoginUserErrorGeneratingJWT

***VerifyEmail:
TestVerifyEmail
TestVerifyEmailErrorInvalidCode
TestVerifyEmailErrorAlreadyVerified

***ResetPassword:
TestResetPassword
TestResetPasswordErrorInvalidEmail
TestResetPasswordErrorUpdatingUser

***ResetPasswordConfirm:
TestResetPasswordConfirm
TestResetPasswordConfirmErrorPasswordsDontMatch
TestResetPasswordConfirmErrorHashingPassword
TestResetPasswordConfirmErrorInvalidCode
TestResetPasswordConfirmErrorUpdatingUser
*/
package authService

import (
	"errors"
	userClient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	authUtils "mvc-go/utils/auth"
	"mvc-go/utils/email"
	"testing"
	"time"

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

// Tests for RegisterUser
func TestRegisterUser(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockEmailClient := email.EmailClient.(*email.EmailMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)

	registerDto := dto.Register{
		FirstName:       "Nombre",
		LastName:        "Apellido",
		Email:           "test@mail.com",
		UserName:        "testname",
		Password:        "pass123",
		PasswordConfirm: "pass123",
	}
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "LApellido",
		Email:            "test@mail.com",
		UserName:         "testname",
		Password:         "hashedpassword",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         false,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockPassClient.On("HashPassword", registerDto.Password).Return(userModel.Password, nil)
	mockClient.On("InsertUser", mock.Anything).Return(userModel)
	mockEmailClient.On("SendEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()

	registerDtoRes, err := AuthService.RegisterUser(registerDto)

	assert.Equal(t, registerDto.FirstName, registerDtoRes.FirstName)
	assert.NotEqual(t, "", registerDtoRes.VerificationCode)
	assert.Nil(t, err)
}
func TestRegisterUserErrorPasswordsDontMatch(t *testing.T) {
	initTestClient()

	registerDto := dto.Register{
		FirstName:       "Nombre",
		LastName:        "Apellido",
		Email:           "test@mail.com",
		UserName:        "testname",
		Password:        "pass123",
		PasswordConfirm: "otraPass",
	}

	registerDtoRes, err := AuthService.RegisterUser(registerDto)
	assert.Equal(t, "", registerDtoRes.FirstName)
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
	assert.Equal(t, "Passwords do not match", err.Message())
}
func TestRegisterUserErrorHashingPassword(t *testing.T) {
	initTestClient()
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)

	registerDto := dto.Register{
		FirstName:       "Nombre",
		LastName:        "Apellido",
		Email:           "test@mail.com",
		UserName:        "testname",
		Password:        "pass123",
		PasswordConfirm: "pass123",
	}

	mockPassClient.On("HashPassword", registerDto.Password).Return("", errors.New(""))

	registerDtoRes, err := AuthService.RegisterUser(registerDto)

	assert.Equal(t, "", registerDtoRes.FirstName)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error hashing password", err.Message())
}
func TestRegisterUserErrorUserAlreadyExists(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)

	registerDto := dto.Register{
		FirstName:       "Nombre",
		LastName:        "Apellido",
		Email:           "test@mail.com",
		UserName:        "testname",
		Password:        "pass123",
		PasswordConfirm: "pass123",
	}
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
		Verified:         false,
		CreatedAt:        time_zero,
		UpdatedAt:        time_zero,
	}

	mockPassClient.On("HashPassword", registerDto.Password).Return(userModel.Password, nil)
	mockClient.On("InsertUser", mock.Anything).Return(userModel)

	registerDtoRes, err := AuthService.RegisterUser(registerDto)

	assert.Equal(t, "", registerDtoRes.FirstName)
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
	assert.Equal(t, "Email or UserName already exists", err.Message())
}

// Tests for LoginUser
func TestLoginUserWithEmail(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)

	var email string = "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockClient.On("GetUserByEmail", email).Return(userModel)
	loginDto := dto.Login{
		Email:    email,
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(nil)
	mockTokenClient.On("GenerateToken", mock.Anything, mock.Anything, mock.Anything).Return("aJWT", nil)

	token, err := AuthService.LoginUser(loginDto)

	assert.Nil(t, err)
	assert.NotEqual(t, "", token)
}
func TestLoginUserWithUserName(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)

	var email string = "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockClient.On("GetUserByUserName", "testname").Return(userModel)
	loginDto := dto.Login{
		UserName: "testname",
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(nil)
	mockTokenClient.On("GenerateToken", mock.Anything, mock.Anything, mock.Anything).Return("aJWT", nil)

	token, err := AuthService.LoginUser(loginDto)

	assert.Nil(t, err)
	assert.NotEqual(t, "", token)
}
func TestLoginUserErrorUserRequired(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)

	var email string = "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockClient.On("GetUserByUserName", "testname").Return(userModel)
	loginDto := dto.Login{
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(nil)
	mockTokenClient.On("GenerateToken", mock.Anything, mock.Anything, mock.Anything).Return("aJWT", nil)

	token, err := AuthService.LoginUser(loginDto)

	assert.NotNil(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, 400, err.Status())
	assert.Equal(t, "Email or user name required", err.Message())
}
func TestLoginUserErrorNotVerified(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
	var email string = "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         false,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockClient.On("GetUserByEmail", email).Return(userModel)
	loginDto := dto.Login{
		Email:    email,
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(nil)
	token, err := AuthService.LoginUser(loginDto)

	assert.NotNil(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, 403, err.Status())
	assert.Equal(t, "Please verify your email", err.Message())
}
func TestLoginUserErrorInvalidEmail(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
	var email string = "email@test.com"
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
		Verified:         false,
		Active:           true,
		CreatedAt:        time_zero,
		UpdatedAt:        time_zero,
	}

	mockClient.On("GetUserByEmail", email).Return(userModel)
	loginDto := dto.Login{
		Email:    email,
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(nil)

	token, err := AuthService.LoginUser(loginDto)

	assert.NotNil(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, 401, err.Status())
	assert.Equal(t, "Invalid email or Password", err.Message())
}
func TestLoginUserErrorInvalidUserName(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
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
		Verified:         false,
		Active:           true,
		CreatedAt:        time_zero,
		UpdatedAt:        time_zero,
	}

	mockClient.On("GetUserByUserName", "testusername").Return(userModel)
	loginDto := dto.Login{
		UserName: "testUserName",
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(nil)

	token, err := AuthService.LoginUser(loginDto)

	assert.NotNil(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, 401, err.Status())
	assert.Equal(t, "Invalid user name or Password", err.Message())
}
func TestLoginUserErrorInvalidPassword(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
	var email string = "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockClient.On("GetUserByEmail", email).Return(userModel)
	loginDto := dto.Login{
		Email:    email,
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(errors.New(""))

	token, err := AuthService.LoginUser(loginDto)

	assert.NotNil(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, 401, err.Status())
	assert.Equal(t, "Invalid email or Password", err.Message())
}
func TestLoginUserErrorGeneratingJWT(t *testing.T) {
	initTestClient()

	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)
	var email string = "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "passwordHashed",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		Active:           true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockClient.On("GetUserByEmail", email).Return(userModel)
	loginDto := dto.Login{
		Email:    email,
		Password: "passwordPayload",
	}
	mockPassClient.On("VerifyPassword", userModel.Password, loginDto.Password).Return(nil)
	mockTokenClient.On("GenerateToken", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	token, err := AuthService.LoginUser(loginDto)

	assert.NotNil(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error generating token", err.Message())
}

func TestRefresh(t *testing.T) {
	initTestClient()
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)

	user_id := uuid.New()
	mockTokenClient.On("GenerateToken", mock.Anything, user_id.String(), mock.Anything).Return("aJWT", nil)

	token, err := AuthService.Refresh(user_id)

	assert.Nil(t, err)
	assert.NotEqual(t, "", token)
	assert.Equal(t, "aJWT", token)
}
func TestRefreshErrorGeneratingToken(t *testing.T) {
	initTestClient()
	mockTokenClient := authUtils.TokenClient.(*authUtils.TokenMockClient)

	user_id := uuid.New()
	mockTokenClient.On("GenerateToken", mock.Anything, user_id.String(), mock.Anything).Return("", errors.New(""))

	token, err := AuthService.Refresh(user_id)

	assert.NotNil(t, err)
	assert.Equal(t, "", token)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error generating token", err.Message())
}

// Tests for VerifyEmail
func TestVerifyEmail(t *testing.T) {
	initTestClient()
	mockclient := userClient.UserClient.(*userClient.UserMockClient)
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
		Verified:         false,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockclient.On("GetUserByCode", mock.AnythingOfType("string")).Return(userModel)

	updatedUser := userModel
	updatedUser.VerificationCode = ""
	updatedUser.Verified = true
	mockclient.On("UpdateUser", updatedUser).Return(updatedUser)

	err := AuthService.VerifyEmail("codeExample")

	assert.Nil(t, err)
}
func TestVerifyEmailErrorInvalidCode(t *testing.T) {
	initTestClient()
	mockclient := userClient.UserClient.(*userClient.UserMockClient)
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
		Verified:         false,
		CreatedAt:        time_zero,
		UpdatedAt:        time_zero,
	}
	mockclient.On("GetUserByCode", mock.AnythingOfType("string")).Return(userModel)

	err := AuthService.VerifyEmail("codeExample")

	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "Invalid verification code or user doesn't exists", err.Message())
}
func TestVerifyEmailErrorAlreadyVerified(t *testing.T) {
	initTestClient()
	mockclient := userClient.UserClient.(*userClient.UserMockClient)
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
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockclient.On("GetUserByCode", mock.AnythingOfType("string")).Return(userModel)

	err := AuthService.VerifyEmail("codeExample")

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
	assert.Equal(t, "User already verified", err.Message())
}

// Tests for ResetPassword
func TestResetPassword(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockEmailClient := email.EmailClient.(*email.EmailMockClient)

	email := "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "hashedpassword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockClient.On("GetUserByEmail", email).Return(userModel)
	updatedUser := userModel
	updatedUser.VerificationCode = "new-code"
	mockClient.On("UpdateUser", mock.Anything).Return(updatedUser)
	mockEmailClient.On("SendEmail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()

	resetPassDto, err := AuthService.ResetPassword(dto.ResetPassword{Email: email})

	assert.Nil(t, err)
	assert.Equal(t, email, resetPassDto.Email)
}
func TestResetPasswordErrorInvalidEmail(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)

	email := "email@test.com"
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
		Verified:         false,
		CreatedAt:        time_zero,
		UpdatedAt:        time_zero,
	}
	mockClient.On("GetUserByEmail", email).Return(userModel)

	resetPassDto, err := AuthService.ResetPassword(dto.ResetPassword{Email: email})

	assert.NotNil(t, err)
	assert.Equal(t, dto.ResetPassword{}, resetPassDto)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "Invalid email", err.Message())
}
func TestResetPasswordErrorUpdatingUser(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)

	email := "email@test.com"
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "Apellido",
		Email:            email,
		UserName:         "testname",
		Password:         "hashedpassword",
		Role:             "user",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	mockClient.On("GetUserByEmail", email).Return(userModel)
	mockClient.On("UpdateUser", mock.Anything).Return(model.User{})

	resetPassDto, err := AuthService.ResetPassword(dto.ResetPassword{Email: email})

	assert.NotNil(t, err)
	assert.Equal(t, dto.ResetPassword{}, resetPassDto)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error generating reset code", err.Message())
}

// Tests for ResetPasswordConfirm
func TestResetPasswordConfirm(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)

	resetPassConfirmDto := dto.ResetPasswordConfirm{
		Password:         "pass123",
		PasswordConfirm:  "pass123",
		VerificationCode: "aexamplecode",
	}
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "LApellido",
		Email:            "test@mail.com",
		UserName:         "testname",
		Password:         "hashedpassword",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockPassClient.On("HashPassword", resetPassConfirmDto.Password).Return(userModel.Password, nil)
	mockClient.On("GetUserByCode", mock.Anything).Return(userModel)
	mockClient.On("UpdateUser", mock.Anything, mock.Anything).Return(userModel)

	err := AuthService.ResetPasswordConfirm(resetPassConfirmDto)

	assert.Nil(t, err)
}
func TestResetPasswordConfirmErrorPasswordsDontMatch(t *testing.T) {
	initTestClient()

	resetPassConfirmDto := dto.ResetPasswordConfirm{
		Password:         "pass123",
		PasswordConfirm:  "otherPass",
		VerificationCode: "aexamplecode",
	}

	err := AuthService.ResetPasswordConfirm(resetPassConfirmDto)

	assert.NotNil(t, err)
	assert.Equal(t, 400, err.Status())
	assert.Equal(t, "Passwords do not match", err.Message())
}
func TestResetPasswordConfirmErrorHashingPassword(t *testing.T) {
	initTestClient()
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)

	resetPassConfirmDto := dto.ResetPasswordConfirm{
		Password:         "pass123",
		PasswordConfirm:  "pass123",
		VerificationCode: "aexamplecode",
	}

	mockPassClient.On("HashPassword", resetPassConfirmDto.Password).Return("", errors.New(""))

	err := AuthService.ResetPasswordConfirm(resetPassConfirmDto)

	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error hashing password", err.Message())
}
func TestResetPasswordConfirmErrorInvalidCode(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)

	resetPassConfirmDto := dto.ResetPasswordConfirm{
		Password:         "pass123",
		PasswordConfirm:  "pass123",
		VerificationCode: "aexamplecode",
	}

	mockPassClient.On("HashPassword", resetPassConfirmDto.Password).Return("aHashedPass", nil)
	mockClient.On("GetUserByCode", mock.Anything).Return(model.User{})

	err := AuthService.ResetPasswordConfirm(resetPassConfirmDto)

	assert.NotNil(t, err)
	assert.Equal(t, 404, err.Status())
	assert.Equal(t, "Invalid verification code", err.Message())
}
func TestResetPasswordConfirmErrorUpdatingUser(t *testing.T) {
	initTestClient()
	mockClient := userClient.UserClient.(*userClient.UserMockClient)
	mockPassClient := authUtils.PassClient.(*authUtils.PassMockClient)

	resetPassConfirmDto := dto.ResetPasswordConfirm{
		Password:         "pass123",
		PasswordConfirm:  "pass123",
		VerificationCode: "aexamplecode",
	}
	user_id := uuid.New()
	now := time.Now()
	userModel := model.User{
		UserID:           user_id,
		FirstName:        "Nombre",
		LastName:         "LApellido",
		Email:            "test@mail.com",
		UserName:         "testname",
		Password:         "hashedpassword",
		Role:             "user",
		VerificationCode: "codeExample",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	mockPassClient.On("HashPassword", resetPassConfirmDto.Password).Return(userModel.Password, nil)
	mockClient.On("GetUserByCode", mock.Anything).Return(userModel)
	mockClient.On("UpdateUser", mock.Anything, mock.Anything).Return(model.User{})

	err := AuthService.ResetPasswordConfirm(resetPassConfirmDto)

	assert.NotNil(t, err)
	assert.Equal(t, 500, err.Status())
	assert.Equal(t, "Error trying reset password", err.Message())
}
