package broker

import "log"

// InitConsumer connects to RabbitMQ, declares and consumes queues
func InitConsumer() {
	Connect()

	declareEmailQueue()
	declareAPNPushQueue()
	declareFCMPushQueue()

	forever := make(chan bool)
	consumeEmailQueue()
	consumeAPNPushQueue()
	consumeFCMPushQueue()

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

func consumeAPNPushQueue() {
	msgs, err := ch.Consume(
		apnPushQueue.Name,
		"APN Push Notifications",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to register APN Push notifications queue consumer:", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Message APN: %s", d.Body)
			d.Ack(false)
		}
	}()
}

func consumeFCMPushQueue() {
	msgs, err := ch.Consume(
		fcmPushQueue.Name,
		"FCM Push Notifications",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to register FCM Push notifications queue consumer:", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Message FCM: %s", d.Body)
			d.Ack(false)
		}
	}()
}
