package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	checkController "rest-chat/src/api/controllers/check"
	messageController "rest-chat/src/api/controllers/message"
	userController "rest-chat/src/api/controllers/user"
)

var (
	routerHandler *mux.Router
)

func main() {
	routerHandler = mux.NewRouter()
	routerHandler.Use(commonMiddleware)

	// Check
	routerHandler.HandleFunc("/check", checkController.Check).Methods(http.MethodGet)

	// User
	routerHandler.HandleFunc("/users", userController.CreateUser).Methods("POST")
	routerHandler.HandleFunc("/login", userController.LoginUser).Methods("POST")
	routerHandler.HandleFunc("/authenticated", userController.AuthenticatedUser).Methods("GET")

	// Messages
	routerHandler.HandleFunc("/messages", messageController.PostMessage).Methods("POST")
	routerHandler.HandleFunc("/messages", messageController.GetMessages).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", routerHandler))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
