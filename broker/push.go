package broker

import (
	"log"

	"github.com/streadway/amqp"
)

var pushQueue *amqp.Queue

func declarePushQueue() {
	var err error

	if pushQueue, err = declareQueue("notifications.push"); err != nil {
		log.Fatalln("Failed to declare Push notifications queue:", err)
	}
}

// PublishToPushQueue publish a message to the push notification queue
func PublishToPushQueue(body []byte) {
	publishToQueue(
		pushQueue,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		},
	)
}
