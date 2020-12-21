package models

// NotificationType is the notification type used inside the Notification struct
type NotificationType string

const (
	// Email notification type
	Email NotificationType = "email"
	// Push notification type
	Push NotificationType = "push"
)

// Notification structure used inside the POST method handler
type Notification struct {
	Type          []NotificationType `json:"type" example:"[\"email\", \"push\"]"`
	Title         string             `json:"title" example:"My Notification"`
	Message       string             `json:"message" example:"Hello my friend"`
	ExpoPushToken string             `json:"expo_push_token" example:"ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]"`
	Email         string             `json:"email" example:"john.doe@foo.com"`
}

// EmailNotification structure used when publishing a new message inside the "email notifications" RabbitMQ queue
type EmailNotification struct {
	Type    NotificationType `json:"type" example:"email"`
	Title   string           `json:"title" example:"My Notification"`
	Message string           `json:"message" example:"Hello my friend"`
	Email   string           `json:"email" example:"john.doe@foo.com"`
}

// PushNotification structure used when publishing a new message inside one of the "push notifications" RabbitMQ queue
type PushNotification struct {
	Type          NotificationType `json:"type" example:"push"`
	Title         string           `json:"title" example:"My Notification"`
	Message       string           `json:"message" example:"Hello my friend"`
	ExpoPushToken string           `json:"expo_push_token" example:"ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]"`
}
