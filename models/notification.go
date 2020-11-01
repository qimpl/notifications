package models

// NotificationType is the notification type used inside the Notification struct
type NotificationType string

const (
	// Email notification type
	Email NotificationType = "email"
	// APN Push notification type
	APN NotificationType = "apn"
	// FCM Push notification type
	FCM NotificationType = "fcm"
)

// Notification structure used inside the POST method handler
type Notification struct {
	Type     []NotificationType `json:"type" example:"[\"email\", \"apn\"]"`
	Title    string             `json:"title" example:"My Notification"`
	Message  string             `json:"message" example:"Hello my friend"`
	DeviceID string             `json:"device_id" example:"00000-00000-00000-00000"`
	Email    string             `json:"email" example:"john.doe@foo.com"`
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
	Type     NotificationType `json:"type" example:"apn"`
	Title    string           `json:"title" example:"My Notification"`
	Message  string           `json:"message" example:"Hello my friend"`
	DeviceID string           `json:"device_id" example:"00000-00000-00000-00000"`
}
