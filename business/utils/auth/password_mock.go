package authUtils

import (
	"github.com/stretchr/testify/mock"
)

type PassMockClient struct {
	mock.Mock
}

func (c *PassMockClient) HashPassword(password string) (string, error) {
	ret := c.Called(password)
	return ret.String(0), ret.Error(1)
}

func (c *PassMockClient) VerifyPassword(hashedPassword string, candidatePassword string) error {
	ret := c.Called(hashedPassword, candidatePassword)
	return ret.Error(0)
}
