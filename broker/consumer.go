package broker

import (
	"encoding/json"
	"log"

	"github.com/qimpl/notifications/models"
	"github.com/qimpl/notifications/services"
)

// InitConsumer connects to RabbitMQ, declares and consumes queues
func InitConsumer() {
	Connect()

	declareEmailQueue()
	declarePushQueue()

	forever := make(chan bool)
	consumeEmailQueue()
	consumePushQueue()

	log.Println("Waiting messages")
	<-forever
}

func consumeEmailQueue() {
	msgs, err := ch.Consume(
		emailQueue.Name,
		"Email Notifications",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to register Email notifications queue consumer:", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Message Email: %s", d.Body)
			d.Ack(false)
		}
	}()

}

func consumePushQueue() {
	msgs, err := ch.Consume(
		pushQueue.Name,
		"APN Push Notifications",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to register Push notifications queue consumer:", err)
	}

	go func() {
		for d := range msgs {
			notification := models.PushNotification{}
			json.Unmarshal(d.Body, &notification)
			services.PushNotification(&notification)
			d.Ack(false)
		}
	}()
}
