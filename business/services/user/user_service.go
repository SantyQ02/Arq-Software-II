package userService

import (
	"errors"
	userClient "mvc-go/clients/user"
	"mvc-go/dto"
	"time"

	e "mvc-go/utils/errors"

	"github.com/google/uuid"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id uuid.UUID) (dto.UserResponse, e.ApiError)
	GetUsers() (dto.UserResponses, e.ApiError)
	DeleteUser(id uuid.UUID) e.ApiError
	MakeAdminUser(id uuid.UUID) (dto.UserResponse, e.ApiError)
	UpdateUser(userDto dto.UserResponse) (dto.UserResponse, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id uuid.UUID) (dto.UserResponse, e.ApiError) {
	idString := id.String()
	user := userClient.UserClient.GetUserById(idString)
	if user.UserID == uuid.Nil {
		return dto.UserResponse{}, e.NewNotFoundApiError("User not found")

	}

	userDto := dto.UserResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Active:    user.Active,
	}
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UserResponses, e.ApiError) {
	users := userClient.UserClient.GetUsers()
	if len(users) == 0 {
		return dto.UserResponses{}, e.NewInternalServerApiError("Error getting users from database", errors.New("Error in database"))
	}

	var usersDto []dto.UserResponse

	for _, user := range users {
		var userDto dto.UserResponse
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.UserID = user.UserID
		userDto.CreatedAt = user.CreatedAt
		userDto.UpdatedAt = user.UpdatedAt
		userDto.Email = user.Email
		userDto.Role = user.Role
		userDto.Active = user.Active

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) DeleteUser(id uuid.UUID) e.ApiError {
	idString := id.String()

	err := userClient.UserClient.DeleteUser(idString)
	if err != nil {
		return e.NewInternalServerApiError("Something went wrong deleting user", nil)
	}

	return nil
}

func (s *userService) MakeAdminUser(id uuid.UUID) (dto.UserResponse, e.ApiError) {
	idString := id.String()
	user := userClient.UserClient.GetUserById(idString)
	if user.UserID == uuid.Nil {
		return dto.UserResponse{}, e.NewNotFoundApiError("User not found")

	}
	user.Role = "admin"
	user.UpdatedAt = time.Now()
	user = userClient.UserClient.UpdateUser(user)
	if user.UserID == uuid.Nil {
		return dto.UserResponse{}, e.NewInternalServerApiError("Error updating user", nil)

	}

	userDto := dto.UserResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Active:    user.Active,
	}
	return userDto, nil
}

func (s *userService) UpdateUser(userDto dto.UserResponse) (dto.UserResponse, e.ApiError) {
	idString := userDto.UserID.String()
	user := userClient.UserClient.GetUserById(idString)
	if user.UserID == uuid.Nil {
		return dto.UserResponse{}, e.NewNotFoundApiError("User not found")

	}
	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Role = userDto.Role
	user.UpdatedAt = time.Now()
	user.Active = userDto.Active

	user = userClient.UserClient.UpdateUser(user)
	if user.UserID == uuid.Nil {
		return dto.UserResponse{}, e.NewInternalServerApiError("Error updating user", nil)

	}

	userDto = dto.UserResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Active:    user.Active,
	}
	return userDto, nil
}
