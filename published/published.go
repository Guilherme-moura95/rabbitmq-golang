package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	// Estabelece conexão com o servidor RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Erro ao estabelecer conexão com o servidor RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Cria um canal de comunicação
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Erro ao criar canal de comunicação: %s", err)
	}
	defer ch.Close()

	// Declara uma fila
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
	for i := 0; i < 10000; i++ {

		// Publica uma mensagem na fila
		message := fmt.Sprintf("Hello, world! %d", i)

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err != nil {
			log.Fatalf("Erro ao publicar mensagem na fila: %s", err)
			log.Printf("Mensagem publicada: %s %i\n", message, i)
		}
	}

}
