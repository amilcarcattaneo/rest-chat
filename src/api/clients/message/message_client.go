package message

import (
	"encoding/json"
	"time"

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
	errMessageNotFound = "Message %s not found"
)

// GetMessages returns all the messages found
func GetMessages(sender, recipient, start, limit uint) ([]domain.Message, error) {
	var messages []domain.Message
	if err := Client.Where("sender = ? AND recipient = ? AND id >= ?", sender, recipient, start).Limit(limit).Find(&messages).Error; err != nil {
		return nil, err
	}
	for i := range messages {
		var content domain.Content
		json.Unmarshal([]byte(messages[i].ContentString), &content)
		messages[i].Content = content
	}
	return messages, nil
}

// PutMessage creates a new message
func PutMessage(now time.Time, message domain.Message) (uint, time.Time, error) {
	content, err := json.Marshal(message.Content)
	if err != nil {
		return 0, time.Time{}, err
	}
	message.ContentString = string(content)
	message.Timestamp = now
	if err := Client.Create(&message).Error; err != nil {
		return 0, time.Time{}, err
	}
	return message.ID, message.Timestamp, nil
}
