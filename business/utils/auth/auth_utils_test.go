package authUtils

import (
	"encoding/json"
	"mvc-go/dto"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func initTestClient() {
	PassClient = &passClient{}
	TokenClient = &tokenClient{}
}

// Tests for password.go
func TestHashPassword(t *testing.T) {
	initTestClient()
	pass := "una password"
	hashedPass, err := PassClient.HashPassword(pass)

	assert.NotEqual(t, hashedPass, pass)
	assert.Nil(t, err)
}
func TestVerifyPassword(t *testing.T) {
	initTestClient()
	pass := "una password"
	hashedPass, err := PassClient.HashPassword(pass)

	err = PassClient.VerifyPassword(hashedPass, pass)

	assert.Nil(t, err)
}
func TestVerifyPasswordErrorDifferentPasswords(t *testing.T) {
	initTestClient()
	pass := "una password"
	hashedPass, err := PassClient.HashPassword(pass)
	if err != nil {
		return
	}

	differentPass := "12345"
	err = PassClient.VerifyPassword(hashedPass, differentPass)

	assert.NotNil(t, err)
}

// Tests for encode.go
func TestEncode(t *testing.T) {
	initTestClient()
	code := "a string 123123"
	encoded := Encode(code)

	assert.NotEqual(t, code, encoded)
	assert.NotEqual(t, "", encoded)
}
func TestDecode(t *testing.T) {
	initTestClient()
	code := "a code 3242"
	decoded, err := Decode(Encode(code))

	assert.Equal(t, code, decoded)
	assert.Nil(t, err)
}

// Test for token.go
func TestGenerateToken(t *testing.T) {
	initTestClient()
	tokenDuration, _ := time.ParseDuration("60m")
	userID := uuid.New()
	_, err := TokenClient.GenerateToken(tokenDuration, userID.String(), "aSecretKey")

	assert.Nil(t, err)
}
func TestValidateToken(t *testing.T) {
	initTestClient()
	tokenDuration, _ := time.ParseDuration("60m")
	userID := uuid.New()
	token, err := TokenClient.GenerateToken(tokenDuration, userID.String(), "aSecretKey")

	userValid, err := TokenClient.ValidateToken(token, "aSecretKey")

	assert.Nil(t, err)
	assert.Equal(t, userID.String(), userValid)
}
func TestValidateTokenError(t *testing.T) {
	initTestClient()
	tokenDuration, _ := time.ParseDuration("60m")
	userDto := dto.UserResponse{FirstName: "minombre"}
	userDtoJSON, _ := json.Marshal(userDto)
	token, err := TokenClient.GenerateToken(tokenDuration, string(userDtoJSON), "aSecretKey")

	userValid, err := TokenClient.ValidateToken(token, "anotherKey")

	assert.NotNil(t, err)
	assert.Equal(t, nil, userValid)
}
