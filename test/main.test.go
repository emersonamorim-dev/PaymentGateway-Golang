package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"paymentgateway-golang/pkg/payments/credit"
	"paymentgateway-golang/pkg/payments/debit"
	"paymentgateway-golang/pkg/payments/pix"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProcessDebitCardPayment(t *testing.T) {
	// Criação do contexto de teste
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/payment/debit", processDebitCardPayment)

	// Dados de pagamento para teste
	payment := debit.DebitPayment{
		CardNumber: "1234567890123456",
		Amount:     100,
	}

	// Converte os dados de pagamento para JSON
	jsonData, _ := json.Marshal(payment)

	// Simula uma requisição HTTP POST para a rota
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payment/debit", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Verifica o status code esperado
	if w.Code != 202 {
		t.Errorf("Esperado código 202, mas obteve %d", w.Code)
	}

	// Verifica a resposta esperada
	expectedResponse := `{"message":"O pagamento com cartão de débito está sendo processado"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Resposta esperada %s, mas obteve %s", expectedResponse, w.Body.String())
	}
}

func TestProcessCreditCardPayment(t *testing.T) {
	// Criação do contexto de teste
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/payment/credit", processCreditCardPayment)

	// Dados de pagamento para teste
	payment := credit.CreditPayment{
		CardNumber: "1234567890123456",
		Amount:     100,
	}

	// Converte os dados de pagamento para JSON
	jsonData, _ := json.Marshal(payment)

	// Simula uma requisição HTTP POST para a rota
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payment/credit", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Verifica o status code esperado
	if w.Code != 202 {
		t.Errorf("Esperado código 202, mas obteve %d", w.Code)
	}

	// Verifica a resposta esperada
	expectedResponse := `{"message":"O pagamento com cartão de crédito está sendo processado"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Resposta esperada %s, mas obteve %s", expectedResponse, w.Body.String())
	}
}

func TestProcessPixPayment(t *testing.T) {
	// Criação do contexto de teste
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/payment/pix", processPixPayment)

	// Dados de pagamento para teste
	payment := pix.PixPayment{
		PixKey:  "chave_pix",
		Amount:  100,
	}

	// Converte os dados de pagamento para JSON
	jsonData, _ := json.Marshal(payment)

	// Simula uma requisição HTTP POST para a rota
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payment/pix", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Verifica o status code esperado
	if w.Code != 202 {
		t.Errorf("Esperado código 202, mas obteve %d", w.Code)
	}

	// Verifica a resposta esperada
	expectedResponse := `{"message":"Pagamento com Pix está sendo processado"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("Resposta esperada %s, mas obteve %s", expectedResponse, w.Body.String())
	}
}
