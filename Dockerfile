FROM golang:1.16

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Exponha a porta 8081
EXPOSE 8081

# Comando para iniciar o aplicativo
CMD ["./main"]
