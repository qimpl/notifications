package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/qimpl/notifications/broker"
	"github.com/qimpl/notifications/models"
)

// AddNotification add a new notification into the broker queues
// @Summary Add notification
// @Description Add a new notification into the broker queues
// @Tags Notifications
// @Accept json
// @Param Notification body models.Notification true "Notification object"
// @Produce json
// @Success 204 ""
// @Failure 422 {string} models.ErrorResponse
// @Router /notification [post]
func AddNotification(w http.ResponseWriter, r *http.Request) {
	var notification models.Notification

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&notification); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		var unprocessableEntity *models.UnprocessableEntity
		json.NewEncoder(w).Encode(unprocessableEntity.GetError("Malformed body"))

		return
	}

	for _, notificationType := range notification.Type {
		switch notificationType {
		case models.Email:
			emailNotification, _ := json.Marshal(models.EmailNotification{
				Type:    models.Email,
				Title:   notification.Title,
				Message: notification.Message,
				Email:   notification.Email,
			})

			broker.PublishToEmailQueue(emailNotification)
		case models.APN:
			APNNotification, _ := json.Marshal(models.PushNotification{
				Type:     models.APN,
				Title:    notification.Title,
				Message:  notification.Message,
				DeviceID: notification.DeviceID,
			})

			broker.PublishToAPNPushQueue(APNNotification)
		case models.FCM:
			FCMNotification, _ := json.Marshal(models.PushNotification{
				Type:     models.FCM,
				Title:    notification.Title,
				Message:  notification.Message,
				DeviceID: notification.DeviceID,
			})

			broker.PublishToFCMPushQueue(FCMNotification)
		default:
			continue
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
