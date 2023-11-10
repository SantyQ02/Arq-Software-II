package authService

import (
	"errors"
	userClient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	"mvc-go/utils/email"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	authUtils "mvc-go/utils/auth"
	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type authService struct{}

type authServiceInterface interface {
	RegisterUser(registerDto dto.Register) (dto.Register, e.ApiError)
	LoginUser(loginDto dto.Login) (string, e.ApiError)
	VerifyEmail(code string) e.ApiError
	Refresh(id uuid.UUID) (string, e.ApiError)
	ResetPassword(dtoResetPass dto.ResetPassword) (dto.ResetPassword, e.ApiError)
	ResetPasswordConfirm(resetPassConfirmDto dto.ResetPasswordConfirm) e.ApiError
}

var (
	AuthService authServiceInterface
)

func init() {
	AuthService = &authService{}
}

func (s *authService) RegisterUser(registerDto dto.Register) (dto.Register, e.ApiError) {

	if registerDto.Password != registerDto.PasswordConfirm {
		return dto.Register{}, e.NewBadRequestApiError("Passwords do not match")
	}

	hashedPassword, err := authUtils.PassClient.HashPassword(registerDto.Password)
	if err != nil {
		return dto.Register{}, e.NewInternalServerApiError("Error hashing password", err)

	}

	code := randstr.String(20)

	verification_code := authUtils.Encode(code)
	now := time.Now()
	user := model.User{
		FirstName:        registerDto.FirstName,
		LastName:         registerDto.LastName,
		Email:            strings.ToLower(registerDto.Email),
		UserName:         strings.ToLower(registerDto.UserName),
		Password:         hashedPassword,
		Role:             "user",
		Active:           true,
		VerificationCode: verification_code,
		Verified:         false,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	user = userClient.UserClient.InsertUser(user)
	if user.UserID == uuid.Nil {
		return dto.Register{}, e.NewBadRequestApiError("Email or UserName already exists")

	}
	registerDto.VerificationCode = user.VerificationCode

	var firstName = registerDto.FirstName

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := email.EmailData{
		URL:       os.Getenv("CLIENT_ORIGIN") + "/auth/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	log.Debug("Email debug: url to verify your account: ", emailData.URL)
	email.EmailClient.SendEmail(registerDto.Email, &emailData, "templates/verificateAccount", "verificateAccountLink.html")

	return registerDto, nil
}

func (s *authService) VerifyEmail(code string) e.ApiError {
	verification_code := authUtils.Encode(code)

	var user model.User = userClient.UserClient.GetUserByCode(verification_code)
	if user.UserID == uuid.Nil {
		return e.NewNotFoundApiError("Invalid verification code or user doesn't exists")
	}

	if user.Verified {
		return e.NewBadRequestApiError("User already verified")
	}

	updatedUser := user
	updatedUser.VerificationCode = ""
	updatedUser.Verified = true

	user = userClient.UserClient.UpdateUser(updatedUser)

	return nil
}

func (s *authService) LoginUser(loginDto dto.Login) (string, e.ApiError) {
	if loginDto.Email == "" && loginDto.UserName == "" {
		return "", e.NewBadRequestApiError("Email or user name required")
	}
	var user model.User
	if loginDto.Email != "" {
		user = userClient.UserClient.GetUserByEmail(strings.ToLower(loginDto.Email))
		if user.UserID == uuid.Nil {
			return "", e.NewUnauthorizedApiError("Invalid email or Password")
		}
	} else {
		user = userClient.UserClient.GetUserByUserName(strings.ToLower(loginDto.UserName))
		if user.UserID == uuid.Nil {
			return "", e.NewUnauthorizedApiError("Invalid user name or Password")
		}
	}
	if !user.Active {
		return "", e.NewForbiddenApiError("Your account has been desactivated")
	}

	if !user.Verified {
		return "", e.NewForbiddenApiError("Please verify your email")
	}

	err := authUtils.PassClient.VerifyPassword(user.Password, loginDto.Password)
	if err != nil {
		return "", e.NewUnauthorizedApiError("Invalid email or Password")
	}

	// Generate Token
	tokenDuration, _ := time.ParseDuration(os.Getenv("TOKEN_EXPIRED_IN"))
	token, err := authUtils.TokenClient.GenerateToken(tokenDuration, user.UserID.String(), os.Getenv("TOKEN_SECRET"))
	if err != nil {
		return "", e.NewInternalServerApiError("Error generating token", err)
	}

	return token, nil
}

func (s *authService) Refresh(id uuid.UUID) (string, e.ApiError) {

	tokenDuration, _ := time.ParseDuration(os.Getenv("TOKEN_EXPIRED_IN"))
	token, err := authUtils.TokenClient.GenerateToken(tokenDuration, id.String(), os.Getenv("TOKEN_SECRET"))
	if err != nil {
		return "", e.NewInternalServerApiError("Error generating token", err)
	}

	return token, nil
}

func (s *authService) ResetPassword(dtoResetPass dto.ResetPassword) (dto.ResetPassword, e.ApiError) {
	user := userClient.UserClient.GetUserByEmail(dtoResetPass.Email)
	if user.UserID == uuid.Nil {
		return dto.ResetPassword{}, e.NewNotFoundApiError("Invalid email")
	}

	code := randstr.String(20)

	verification_code := authUtils.Encode(code)
	now := time.Now()
	user.VerificationCode = verification_code
	user.UpdatedAt = now

	user = userClient.UserClient.UpdateUser(user)
	if user.UserID == uuid.Nil {
		return dto.ResetPassword{}, e.NewInternalServerApiError("Error generating reset code", errors.New("Database error"))
	}

	var firstName = user.FirstName

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := email.EmailData{
		URL:       os.Getenv("CLIENT_ORIGIN") + "/auth/reset-pass/" + code,
		FirstName: firstName,
		Subject:   "Your code to reset password",
	}
	log.Debug("Email debug: url to reset your password: ", emailData.URL)

	email.EmailClient.SendEmail(dtoResetPass.Email, &emailData, "templates/resetPassword", "resetPasswordLink.html")
	return dtoResetPass, nil
}

func (s *authService) ResetPasswordConfirm(resetPassConfirmDto dto.ResetPasswordConfirm) e.ApiError {
	if resetPassConfirmDto.Password != resetPassConfirmDto.PasswordConfirm {
		return e.NewBadRequestApiError("Passwords do not match")
	}

	hashedPassword, err := authUtils.PassClient.HashPassword(resetPassConfirmDto.Password)
	if err != nil {
		return e.NewInternalServerApiError("Error hashing password", err)

	}

	verification_code := authUtils.Encode(resetPassConfirmDto.VerificationCode)

	var user model.User = userClient.UserClient.GetUserByCode(verification_code)
	if user.UserID == uuid.Nil {
		return e.NewNotFoundApiError("Invalid verification code")
	}

	user.Password = hashedPassword
	user.VerificationCode = ""
	now := time.Now()
	user.UpdatedAt = now

	user = userClient.UserClient.UpdateUser(user)
	if user.UserID == uuid.Nil {
		return e.NewInternalServerApiError("Error trying reset password", errors.New("Database error"))
	}

	return nil
}
