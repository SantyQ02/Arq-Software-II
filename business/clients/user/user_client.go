package userClient

import (
	"errors"
	"mvc-go/model"
	"strings"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type userClient struct{}

type userClientInterface interface {
	GetUserByCode(code string) model.User
	GetUserById(id string) model.User
	GetUserByEmail(email string) model.User
	GetUserByUserName(userName string) model.User
	GetUsers() model.Users
	InsertUser(user model.User) model.User
	UpdateUser(user model.User) model.User
	DeleteUser(id string) error
}

var (
	UserClient userClientInterface
)

func init() {
	UserClient = &userClient{}
}

var Db *gorm.DB

func (c *userClient) GetUserByCode(code string) model.User {

	var user model.User
	Db.First(&user, "verification_code = ?", code)
	log.Debug("User: ", user)

	return user
}

func (c *userClient) GetUserById(id string) model.User {
	var user model.User

	Db.First(&user, "user_id = ?", id)
	log.Debug("User: ", user)

	return user
}

func (c *userClient) GetUserByEmail(email string) model.User {
	var user model.User
	Db.First(&user, "email = ?", strings.ToLower(email))
	log.Debug("User: ", user)
	return user
}

func (c *userClient) GetUserByUserName(userName string) model.User {
	var user model.User
	Db.First(&user, "user_name = ?", strings.ToLower(userName))
	log.Debug("User: ", user)
	return user
}

func (c *userClient) GetUsers() model.Users {
	var users model.Users
	result := Db.Find(&users)
	if result.Error != nil {
		log.Error("")
		return model.Users{}
	}
	log.Debug("Users: ", users)

	return users
}

func (c *userClient) InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("")
		return model.User{}
	}
	log.Debug("User Created: ", user.UserID)
	return user
}

func (c *userClient) UpdateUser(user model.User) model.User {
	result := Db.Save(&user)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return model.User{}
	}
	return user
}

func (c *userClient) DeleteUser(id string) error {
	var user model.User
	result := Db.Delete(&user, "user_id = ?", id)
	if result.Error != nil {
		log.Debug(id)
		log.Error(result.Error.Error())
		return errors.New(result.Error.Error())
	}
	return nil
}
