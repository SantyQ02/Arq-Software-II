package userClient

import (
	"mvc-go/model"

	"github.com/stretchr/testify/mock"
)

type UserMockClient struct {
	mock.Mock
}

func (c *UserMockClient) GetUserByCode(code string) model.User {
	ret := c.Called(code)
	return ret.Get(0).(model.User)
}

func (c *UserMockClient) GetUserById(id string) model.User {
	ret := c.Called(id)
	return ret.Get(0).(model.User)
}

func (c *UserMockClient) GetUserByEmail(email string) model.User {
	ret := c.Called(email)
	return ret.Get(0).(model.User)
}

func (c *UserMockClient) GetUserByUserName(userName string) model.User {
	ret := c.Called(userName)
	return ret.Get(0).(model.User)
}

func (c *UserMockClient) GetUsers() model.Users {
	ret := c.Called()
	return ret.Get(0).(model.Users)
}

func (c *UserMockClient) InsertUser(user model.User) model.User {
	ret := c.Called(user)
	return ret.Get(0).(model.User)
}

func (c *UserMockClient) UpdateUser(user model.User) model.User {
	ret := c.Called(user)
	return ret.Get(0).(model.User)
}

func (c *UserMockClient) DeleteUser(id string) error {
	ret := c.Called(id)
	return ret.Error(0)
}
