package pix

import (
	"errors"
	"paymentgateway-golang/pkg/database"
	"paymentgateway-golang/pkg/queue"
)

type PixPayment struct {
	PixKey string  `json:"pix_key"`
	Amount float64 `json:"amount"`
}

func ProcessPayment(payment PixPayment) error {
	if payment.PixKey == "" || payment.Amount <= 0 {
		return errors.New("dados de pagamento inválidos")
	}

	// Publicar a mensagem de pagamento no RabbitMQ
	err := queue.PublishPayment(payment)
	if err != nil {
		return err
	}

	// Salvar o pagamento no DynamoDB
	err = database.SavePayment(payment)
	if err != nil {
		return err
	}

	return nil
}
