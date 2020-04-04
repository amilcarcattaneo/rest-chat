package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"rest-chat/src/api/controllers"
	"rest-chat/src/api/domain"
	messageService "rest-chat/src/api/services/message"
)

// GetMessages handler for getting messages
func GetMessages(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	recipient, err := getValueFromParam(r.URL, "recipient", true)
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}
	if recipient == 0 {
		controllers.HandleError(w, http.StatusBadRequest, errors.New("Recipient id is 0"))
		return
	}
	start, err := getValueFromParam(r.URL, "start", true)
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}
	if start == 0 {
		controllers.HandleError(w, http.StatusBadRequest, errors.New("Start is 0"))
		return
	}

	limit, err := getValueFromParam(r.URL, "limit", false)
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}

	messages, err := messageService.GetMessages(tokenString, recipient, start, limit)
	if err != nil && err.Error() == "not authenticated" {
		controllers.HandleError(w, http.StatusUnauthorized, errors.New("not authenticated"))
		return
	}
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}
	messagesOutput := struct {
		Messages []domain.Message `json:"messages"`
	}{
		Messages: messages,
	}
	json.NewEncoder(w).Encode(&messagesOutput)
	w.WriteHeader(http.StatusOK)
}

func getValueFromParam(url *url.URL, paramName string, required bool) (uint, error) {
	queryParams := url.Query()
	valueFromParams, ok := queryParams[paramName]
	if !ok && required {
		return 0, fmt.Errorf("Missing '%s' url param", paramName)
	}
	if !ok {
		return 0, nil
	}
	value, err := strconv.Atoi(valueFromParams[0])
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

// PostMessage handler for posting a message
func PostMessage(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}

	messageID, timestamp, err := messageService.PostMessage(tokenString, message)
	if err != nil && err.Error() == "not authenticated" {
		controllers.HandleError(w, http.StatusUnauthorized, errors.New("not authenticated"))
		return
	}
	if err != nil {
		controllers.HandleError(w, http.StatusBadRequest, err)
		return
	}
	newMessageOutput := struct {
		ID        uint      `json:"id"`
		Timestamp time.Time `json:"timestamp"`
	}{
		ID:        messageID,
		Timestamp: timestamp,
	}
	json.NewEncoder(w).Encode(&newMessageOutput)
	w.WriteHeader(http.StatusOK)
}
