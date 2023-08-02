package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPaymentEndpoint(t *testing.T) {
	r := gin.Default()
	r.POST("/payment", func(c *gin.Context) {
		c.JSON(202, gin.H{
			"message": "O pagamento está sendo processado",
		})
	})

	req, err := http.NewRequest("POST", "/payment", strings.NewReader(`{"amount": 100}`))
	if err != nil {
		t.Fatalf("Não foi possível criar a solicitação: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("Espere para obter status %d but instead got %d", http.StatusAccepted, w.Code)
	}

	if w.Body.String() != `{"message":"O pagamento está sendo processado"}` {
		t.Fatalf("Espere para receber mensagem '%s' but instead got '%s'", `{"message":"O pagamento está sendo processado"}`, w.Body.String())
	}
}
