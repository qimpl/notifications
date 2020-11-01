package broker

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel

// Connect to RabbitMQ and open a channel connection
func Connect() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_ENDPOINT"))

	if err != nil {
		log.Fatalln("RabbitMQ connection failed:", err)
	}

	ch, err = conn.Channel()

	if err != nil {
		log.Fatalln("Channel opening failed:", err)
	}
}
