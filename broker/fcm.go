package broker

import (
	"log"

	"github.com/streadway/amqp"
)

var fcmPushQueue *amqp.Queue

func declareFCMPushQueue() {
	var err error

	if fcmPushQueue, err = declareQueue("notifications.push.fcm"); err != nil {
		log.Fatalln("Failed to declare FCM push notifications queue:", err)
	}
}

// PublishToFCMPushQueue publish a message to the notification FCM queue
func PublishToFCMPushQueue(body []byte) {
	publishToQueue(fcmPushQueue, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
}
