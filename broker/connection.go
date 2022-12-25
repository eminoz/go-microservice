package broker

import (
	"log"

	"github.com/streadway/amqp"
)

func Connect() *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + "eminoz" + ":" + "eminoz" + "@" + "localhost" + ":" + "5672" + "/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	// Use the connection to create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	return ch
}
