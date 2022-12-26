package broker

import (
	"encoding/json"
	"log"

	"github.com/eminoz/go-api/pkg/model"
	"github.com/streadway/amqp"
)

type User interface {
	CreatedUser(u model.User)
}

type userProducer struct{}

func NewUserProducer() User {
	return &userProducer{}
}
func producer() amqp.Queue {
	ch := Connect()
	// Declare a queue
	q, err := ch.QueueDeclare(
		"createdUser", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}
	return q

}
func (u userProducer) CreatedUser(user model.User) {

	ch := GetBrokerConnection()
	q := producer()

	jsonData, err := json.Marshal(user)

	if err != nil {
		log.Fatal("Failed to encode struct:", err)
	}

	ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(jsonData),
		},
	)
}
