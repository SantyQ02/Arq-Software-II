package userClient

import (
	"mvc-go/model"
	"mvc-go/utils/initializers"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	initializers.LoadTestEnv("../../utils/initializers/test.env")
	// DB Connections Paramters
	DBName := os.Getenv("MYSQL_DB_NAME")
	DBUser := os.Getenv("MYSQL_DB_USER")
	DBPass := os.Getenv("MYSQL_DB_PASS")
	DBHost := os.Getenv("MYSQL_DB_HOST")
	DBPort := os.Getenv("MYSQL_DB_PORT")
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":"+DBPort+")/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	Db = db

}

func initTestClient() {
	UserClient = &userClient{}
}

var now = time.Now()
var userID uuid.UUID
var userModel = model.User{
	FirstName:        "Test",
	LastName:         "Test",
	Email:            "emailimposibledecopiar@gmail.com",
	UserName:         "imposibledecopiar",
	Role:             "user",
	Password:         "hashedPassword",
	VerificationCode: "uncodigo",
	Verified:         false,
	CreatedAt:        now,
	UpdatedAt:        now,
}

func TestInsertUser(t *testing.T) {
	initTestClient()
	user := UserClient.InsertUser(userModel)

	userID = user.UserID
	assert.NotEqual(t, uuid.Nil, user.UserID)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, userModel.FirstName, user.FirstName)
	assert.Equal(t, userModel.LastName, user.LastName)
	assert.Equal(t, userModel.UserName, user.UserName)
	assert.Equal(t, userModel.Email, user.Email)
	assert.Equal(t, userModel.Role, user.Role)
	assert.Equal(t, userModel.Password, user.Password)
	assert.Equal(t, userModel.VerificationCode, user.VerificationCode)
	assert.Equal(t, userModel.Verified, user.Verified)
}
func TestInsertUserErrorDuplicate(t *testing.T) {
	initTestClient()
	user := UserClient.InsertUser(userModel)

	assert.Equal(t, uuid.Nil, user.UserID)
	assert.Equal(t, model.User{}, user)
}
func TestGetUserByEmail(t *testing.T) {
	initTestClient()
	user := UserClient.GetUserByEmail(userModel.Email)

	assert.NotEqual(t, uuid.Nil, user.UserID)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, userModel.FirstName, user.FirstName)
	assert.Equal(t, userModel.LastName, user.LastName)
	assert.Equal(t, userModel.UserName, user.UserName)
	assert.Equal(t, userModel.Email, user.Email)
	assert.Equal(t, userModel.Role, user.Role)
	assert.Equal(t, userModel.Password, user.Password)
	assert.Equal(t, userModel.VerificationCode, user.VerificationCode)
	assert.Equal(t, userModel.Verified, user.Verified)
}
func TestGetUserByUserName(t *testing.T) {
	initTestClient()
	user := UserClient.GetUserByUserName(userModel.UserName)

	assert.NotEqual(t, uuid.Nil, user.UserID)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, userModel.FirstName, user.FirstName)
	assert.Equal(t, userModel.LastName, user.LastName)
	assert.Equal(t, userModel.UserName, user.UserName)
	assert.Equal(t, userModel.Email, user.Email)
	assert.Equal(t, userModel.Role, user.Role)
	assert.Equal(t, userModel.Password, user.Password)
	assert.Equal(t, userModel.VerificationCode, user.VerificationCode)
	assert.Equal(t, userModel.Verified, user.Verified)
}
func TestGetUserByCode(t *testing.T) {
	initTestClient()
	user := UserClient.GetUserByCode(userModel.VerificationCode)

	assert.NotEqual(t, uuid.Nil, user.UserID)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, userModel.FirstName, user.FirstName)
	assert.Equal(t, userModel.LastName, user.LastName)
	assert.Equal(t, userModel.UserName, user.UserName)
	assert.Equal(t, userModel.Email, user.Email)
	assert.Equal(t, userModel.Role, user.Role)
	assert.Equal(t, userModel.Password, user.Password)
	assert.Equal(t, userModel.VerificationCode, user.VerificationCode)
	assert.Equal(t, userModel.Verified, user.Verified)
}
func TestGetUserById(t *testing.T) {
	initTestClient()
	user := UserClient.GetUserById(userID.String())

	assert.NotEqual(t, uuid.Nil, user.UserID)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, userModel.FirstName, user.FirstName)
	assert.Equal(t, userModel.LastName, user.LastName)
	assert.Equal(t, userModel.UserName, user.UserName)
	assert.Equal(t, userModel.Email, user.Email)
	assert.Equal(t, userModel.Role, user.Role)
	assert.Equal(t, userModel.Password, user.Password)
	assert.Equal(t, userModel.VerificationCode, user.VerificationCode)
	assert.Equal(t, userModel.Verified, user.Verified)
}
func TestGetUsers(t *testing.T) {
	initTestClient()
	users := UserClient.GetUsers()

	assert.NotEqual(t, 0, len(users))
}
func TestUpdateUser(t *testing.T) {
	initTestClient()
	var newNow = time.Now()
	var userModelUpdate = model.User{
		UserID:           userID,
		FirstName:        "Otro",
		LastName:         "Otro",
		Email:            "emailimposibledecopiar@gmail.com",
		UserName:         "imposibledecopiar2",
		Role:             "admin",
		Password:         "hashedPassword2",
		VerificationCode: "",
		Verified:         true,
		CreatedAt:        now,
		UpdatedAt:        newNow,
	}
	user := UserClient.UpdateUser(userModelUpdate)

	assert.NotEqual(t, uuid.Nil, user.UserID)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, userModelUpdate.FirstName, user.FirstName)
	assert.Equal(t, userModelUpdate.LastName, user.LastName)
	assert.Equal(t, userModelUpdate.UserName, user.UserName)
	assert.Equal(t, userModelUpdate.Email, user.Email)
	assert.Equal(t, userModelUpdate.Role, user.Role)
	assert.Equal(t, userModelUpdate.Password, user.Password)
	assert.Equal(t, userModelUpdate.VerificationCode, user.VerificationCode)
	assert.Equal(t, userModelUpdate.Verified, user.Verified)
}
func TestUpdateUserError(t *testing.T) {
	initTestClient()
	var userModelUpdate = model.User{
		FirstName:        "Test",
		LastName:         "Test",
		Email:            "emailimposibledecopiar@gmail.com",
		UserName:         "imposibledecopiar",
		Role:             "user",
		Password:         "hashedPassword",
		VerificationCode: "uncodigo",
		Verified:         false,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	user := UserClient.UpdateUser(userModelUpdate)

	assert.Equal(t, uuid.Nil, user.UserID)
	assert.Equal(t, model.User{}, user)
}
func TestDeleteUser(t *testing.T) {
	err := UserClient.DeleteUser(userID.String())

	assert.Nil(t, err)
}
