package broker

import (
	"log"

	"github.com/streadway/amqp"
)

var emailQueue *amqp.Queue

func declareEmailQueue() {
	var err error

	if emailQueue, err = declareQueue("notifications.email"); err != nil {
		log.Fatalln("Failed to declare Email notifications queue:", err)
	}
}

// PublishToEmailQueue publish a message to the notification email queue
func PublishToEmailQueue(body []byte) {
	publishToQueue(
		emailQueue,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		},
	)
}
