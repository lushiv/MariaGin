package common_utils

import (
	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

func InitializeRabbitMQConnection(amqpURI string) error {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return err
	}

	RabbitMQConn = conn
	RabbitMQChannel = ch

	return nil
}

func PublishToQueue(queueName string, message []byte) error {
	err := RabbitMQChannel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	return err
}

func ConsumeFromQueue(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := RabbitMQChannel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	return msgs, err
}

func CloseRabbitMQConnection() {
	if RabbitMQChannel != nil {
		RabbitMQChannel.Close()
	}
	if RabbitMQConn != nil {
		RabbitMQConn.Close()
	}
}
