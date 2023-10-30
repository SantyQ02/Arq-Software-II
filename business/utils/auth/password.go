package authUtils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type passClient struct{}

type passClientInterface interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword string, candidatePassword string) error
}

var (
	PassClient passClientInterface
)

func init() {
	PassClient = &passClient{}
}

func (c *passClient) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func (c *passClient) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
