# Payment Gateway em Golang

Codificação em Golang com uso de Framework Gin para desenvolvimento de um Gateway de Pagamento que foi implementado  para lidar com pagamentos com cartão de débito, cartão de crédito e Pix. Uso de Banco de dados DynamoDB da AWS e RabbitMQ para
consumo de filas

## Requisitos

- Go 1.16 ou superior
- RabbitMQ
- AWS DynamoDB

## Instalação

1. Clone este repositório para sua máquina local
2. Navegue até o diretório do projeto
3. Execute `go get` para baixar as dependências necessárias

## Uso

Inicie o servidor executando `go run main.go`. Isso iniciará o servidor na porta 8081.

O servidor aceita as seguintes rotas:

- `POST /payment/debit`: Processa um pagamento com cartão de débito. O corpo da solicitação deve ser um JSON que representa o pagamento.
- `POST /payment/credit`: Processa um pagamento com cartão de crédito. O corpo da solicitação deve ser um JSON que representa o pagamento.
- `POST /payment/pix`: Processa um pagamento com Pix. O corpo da solicitação deve ser um JSON que representa o pagamento.

## Estrutura do Projeto

O projeto segue a estrutura de diretórios padrão do Go:

- `main.go`: O ponto de entrada para o aplicativo. Inicia o servidor e define as rotas.
- `pkg/payments/debit`: Contém a lógica para processar pagamentos com cartão de débito.
- `pkg/payments/credit`: Contém a lógica para processar pagamentos com cartão de crédito.
- `pkg/payments/pix`: Contém a lógica para processar pagamentos com Pix.

## Autor:
Emerso Amorim
