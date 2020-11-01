package broker

import (
	"log"

	"github.com/streadway/amqp"
)

var apnPushQueue *amqp.Queue

func declareAPNPushQueue() {
	var err error

	if apnPushQueue, err = declareQueue("notifications.push.apn"); err != nil {
		log.Fatalln("Failed to declare APN push notifications queue:", err)
	}
}

// PublishToAPNPushQueue publish a message to the notification APN queue
func PublishToAPNPushQueue(body []byte) {
	publishToQueue(apnPushQueue, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
}
