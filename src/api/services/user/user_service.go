package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	userClient "rest-chat/src/api/clients/user"
	"rest-chat/src/api/domain"
)

// CreateUser creates a new user
func CreateUser(newUser domain.User) (uint, error) {

	_, err := userClient.GetUser(newUser.Username)
	if err != nil {
		// user not found => creates a new user
		if err.Error() == userClient.ErrUserNotFound {
			hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 5)
			if err != nil {
				return 0, err
			}

			newUser.Password = string(hash)
			id, err := userClient.CreateUser(newUser)
			if err != nil {
				return 0, err
			}

			return id, nil
		}
		// an actual error happened
		return 0, err
	}
	// User already exists
	return 0, errors.New("Username already exists")
}

// LoginUser login
func LoginUser(userLogin domain.UserLogin) (*domain.LoginResponse, error) {

	var user domain.User
	user, err := userClient.GetUser(userLogin.Username)
	if err != nil {
		return nil, errors.New("invalid user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return nil, errors.New("invalid user")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		ID:    user.ID,
		Token: tokenString,
	}, nil
}

// AuthenticatedUser profile
func AuthenticatedUser(tokenString string) (domain.AuthenticatedResponse, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return domain.AuthenticatedResponse{
			Authenticated: false,
			ID:            0,
		}, err
	}

	var user domain.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.ID = uint(claims["id"].(float64))

		return domain.AuthenticatedResponse{
			Authenticated: true,
			ID:            user.ID,
		}, nil
	}

	return domain.AuthenticatedResponse{
		Authenticated: false,
		ID:            0,
	}, nil
}
