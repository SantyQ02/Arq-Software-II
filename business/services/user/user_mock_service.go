package userService

import (
	"mvc-go/dto"
	e "mvc-go/utils/errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserMockService struct {
	mock.Mock
}

func (s *UserMockService) GetUserById(id uuid.UUID) (dto.UserResponse, e.ApiError) {
	ret := s.Called(id)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.UserResponse), nil
	}
	return ret.Get(0).(dto.UserResponse), ret.Get(1).(e.ApiError)
}
func (s *UserMockService) GetUsers() (dto.UserResponses, e.ApiError) {
	ret := s.Called()
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.UserResponses), nil
	}
	return ret.Get(0).(dto.UserResponses), ret.Get(1).(e.ApiError)
}
func (s *UserMockService) DeleteUser(id uuid.UUID) e.ApiError {
	ret := s.Called(id)
	if ret.Get(0) == nil {
		return nil
	}
	return ret.Get(0).(e.ApiError)
}
func (s *UserMockService) MakeAdminUser(id uuid.UUID) (dto.UserResponse, e.ApiError) {
	ret := s.Called(id)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.UserResponse), nil
	}
	return ret.Get(0).(dto.UserResponse), ret.Get(1).(e.ApiError)
}
func (s *UserMockService) UpdateUser(userDto dto.UserResponse) (dto.UserResponse, e.ApiError) {
	ret := s.Called(userDto)
	if ret.Get(1) == nil {
		return ret.Get(0).(dto.UserResponse), nil
	}
	return ret.Get(0).(dto.UserResponse), ret.Get(1).(e.ApiError)
}
