package debit

import (
	"errors"
	"paymentgateway-golang/pkg/database"
	"paymentgateway-golang/pkg/queue"
)

type DebitPayment struct {
	CardNumber string  `json:"card_number"`
	CVV        string  `json:"cvv"`
	ExpiryDate string  `json:"expiry_date"`
	Amount     float64 `json:"amount"`
}

func ProcessPayment(payment DebitPayment) error {
	if payment.CardNumber == "" || payment.CVV == "" || payment.ExpiryDate == "" || payment.Amount <= 0 {
		return errors.New("invalid payment data")
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
