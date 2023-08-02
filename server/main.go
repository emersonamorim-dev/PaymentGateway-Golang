package main

import (
	"paymentgateway-golang/pkg/payments/credit"
	"paymentgateway-golang/pkg/payments/debit"
	"paymentgateway-golang/pkg/payments/pix"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/payment/debit", processDebitCardPayment)
	r.POST("/payment/credit", processCreditCardPayment)
	r.POST("/payment/pix", processPixPayment)

	r.Run(":8081")
}

func processDebitCardPayment(c *gin.Context) {
	var payment debit.DebitPayment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := debit.ProcessPayment(payment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(202, gin.H{
		"message": "O pagamento com cartão de débito está sendo processado",
	})
}

func processCreditCardPayment(c *gin.Context) {
	var payment credit.CreditPayment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := credit.ProcessPayment(payment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(202, gin.H{
		"message": "O pagamento com cartão de crédito está sendo processado",
	})
}

func processPixPayment(c *gin.Context) {
	var payment pix.PixPayment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := pix.ProcessPayment(payment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(202, gin.H{
		"message": "Pagamento com Pix está sendo processado",
	})
}
