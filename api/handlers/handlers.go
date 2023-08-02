package handlers

import (
	"paymentgateway-golang/pkg/payments/debit"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/payment/debit", processDebitPayment)
}

func processDebitPayment(c *gin.Context) {
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
		"message": "O pagamento está sendo processado",
	})
}
