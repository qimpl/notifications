package broker

import "github.com/streadway/amqp"

// InitProducer connects to RabbitMQ, declares and publish to queues
func InitProducer() {
	Connect()

	declareEmailQueue()
	declarePushQueue()
}

func declareQueue(name string) (*amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &queue, nil
}

func publishToQueue(queue *amqp.Queue, publishing amqp.Publishing) error {
	err := ch.Publish(
		"",
		queue.Name,
		false,
		false,
		publishing,
	)

	if err != nil {
		return err
	}

	return nil
}
