package services

import (
	"fmt"
	"log"

	"github.com/qimpl/notifications/models"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// PushNotification send a push notification to given Expo push tokens
func PushNotification(notification *models.PushNotification) {
	expoPushToken, err := expo.NewExponentPushToken(notification.ExpoPushToken)

	if err != nil {
		log.Println("ERROR", err)
	}

	client := expo.NewPushClient(nil)

	response, err := client.Publish(
		&expo.PushMessage{
			To:       []expo.ExponentPushToken{expoPushToken},
			Body:     notification.Message,
			Data:     map[string]string{"withSome": "data"},
			Sound:    "default",
			Title:    notification.Title,
			Priority: expo.DefaultPriority,
		},
	)

	if err != nil {
		log.Println("PUBLISH ERROR", err)
		return
	}

	if response.ValidateResponse() != nil {
		fmt.Println(response.PushMessage.To, "failed")
	}
}
