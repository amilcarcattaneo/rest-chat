package user

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"rest-chat/src/api/domain"
)

var (
	Client *gorm.DB
)

func init() {
	var err error
	if Client, err = gorm.Open(
		"sqlite3",
		"challenge.db",
	); err != nil {
		panic(err.Error())
	}
}

const (
	// ErrUserNotFound defines
	ErrUserNotFound = "User not found"
)

// GetUser returns the user found
func GetUser(username string) (domain.User, error) {
	var user domain.User
	err := Client.First(&user, domain.User{Username: username}).Error
	if user.Username == "" || err != nil {
		return user, errors.New(ErrUserNotFound)
	}
	return user, nil
}

// CreateUser inserts a new user
func CreateUser(newUser domain.User) (uint, error) {
	if err := Client.Create(&newUser).Error; err != nil {
		return 0, err
	}
	return newUser.ID, nil
}
