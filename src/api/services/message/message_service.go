package message

import (
	"errors"
	"time"

	messageClient "rest-chat/src/api/clients/message"
	"rest-chat/src/api/domain"
	userService "rest-chat/src/api/services/user"
)

// GetMessages returns the messages
func GetMessages(token string, recipient, start, limit uint) ([]domain.Message, error) {
	authenticated, err := userService.AuthenticatedUser(token)
	if err != nil {
		return nil, err
	}
	if !authenticated.Authenticated {
		return nil, errors.New("not authenticated")
	}

	if limit == 0 {
		limit = 100
	}

	return messageClient.GetMessages(authenticated.ID, recipient, start, limit)
}

// PostMessage posts a new message
func PostMessage(token string, message domain.Message) (uint, time.Time, error) {
	authenticated, err := userService.AuthenticatedUser(token)
	if err != nil {
		return 0, time.Time{}, err
	}
	if !authenticated.Authenticated {
		return 0, time.Time{}, errors.New("not authenticated")
	}
	return messageClient.PutMessage(time.Now(), message)
}
