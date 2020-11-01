package router

import (
	"github.com/qimpl/notifications/handlers"

	"github.com/gorilla/mux"
)

func createNotificationRouter(router *mux.Router) {
	notificationRouter := router.PathPrefix("/notification").Subrouter()

	notificationRouter.
		HandleFunc("", handlers.AddNotification).
		Methods("POST")
}
