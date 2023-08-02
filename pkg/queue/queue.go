package queue

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func PublishPayment(payment interface{}) error {
	paymentJson, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Cria um canal
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Publica a mensagem de pagamento no RabbitMQ
	err = ch.Publish(
		"",      // exchange
		"queue", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        paymentJson,
		})
	if err != nil {
		return err
	}

	return nil
}
