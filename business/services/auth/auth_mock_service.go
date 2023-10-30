package authService

import (
	"mvc-go/dto"
	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type AuthMockService struct {
	mock.Mock
}

func (s *AuthMockService) RegisterUser(registerDto dto.Register) (dto.Register, e.ApiError) {
	ret := s.Called(registerDto)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.Register), nil
	}
	return ret.Get(0).(dto.Register), ret.Get(1).(e.ApiError)
}
func (s *AuthMockService) VerifyEmail(code string) e.ApiError {
	ret := s.Called(code)
	if ret.Get(0) == nil {
		return nil
	}
	return ret.Get(0).(e.ApiError)
}
func (s *AuthMockService) LoginUser(loginDto dto.Login) (string, e.ApiError) {
	ret := s.Called(loginDto)
	if ret.Get(1) == nil {
		return ret.String(0), nil
	}
	return ret.String(0), ret.Get(1).(e.ApiError)
}
func (s *AuthMockService) Refresh(id uuid.UUID) (string, e.ApiError) {
	ret := s.Called(id)
	if ret.Get(1) == nil {
		return ret.String(0), nil
	}
	return ret.String(0), ret.Get(1).(e.ApiError)
}
func (s *AuthMockService) ResetPassword(dtoResetPass dto.ResetPassword) (dto.ResetPassword, e.ApiError) {
	ret := s.Called(dtoResetPass)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.ResetPassword), nil
	}
	return ret.Get(0).(dto.ResetPassword), ret.Get(1).(e.ApiError)
}
func (s *AuthMockService) ResetPasswordConfirm(resetPassConfirmDto dto.ResetPasswordConfirm) e.ApiError {
	ret := s.Called(resetPassConfirmDto)
	if ret.Get(0) == nil {
		return nil
	}
	return ret.Get(0).(e.ApiError)
}
