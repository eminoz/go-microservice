package broker

import (
	"bytes"
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
func produce() amqp.Queue {
	ch := Connect()
	// Declare a queue
	q, err := ch.QueueDeclare(
		"createdUser", // name
		false,         // durable
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

	ch := Connect()
	q := produce()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		log.Fatal("Failed to encode struct:", err)
	}

	ch.Publish(
		"usercreated", // exchange
		q.Name,        // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        buf.Bytes(),
		},
	)
}
