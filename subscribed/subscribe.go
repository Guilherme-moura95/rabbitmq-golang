package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	// Connect RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Erro ao estabelecer conexão com o servidor RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Create Channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao criar canal de comunicação: %s", err)
	}
	defer ch.Close()

	// Declare line
	q, err := ch.QueueDeclare(
		"minha_fila",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Erro ao declarar a fila: %s", err)
	}

	// create and consume line
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Erro ao registrar consumidor na fila: %s", err)
	}

	// Recebe mensagens da fila
	for msg := range msgs {
		log.Printf("Mensagem recebida: %s\n", msg.Body)
	}
}
