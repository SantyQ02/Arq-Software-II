package authUtils

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type TokenMockClient struct {
	mock.Mock
}

func (c *TokenMockClient) GenerateToken(ttl time.Duration, payload string, secretJWTKey string) (string, error) {
	ret := c.Called(ttl, payload, secretJWTKey)
	return ret.String(0), ret.Error(1)
}

func (c *TokenMockClient) ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	ret := c.Called(token, signedJWTKey)
	return ret.Get(0), ret.Error(1)
}
