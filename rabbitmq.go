package main

import (
	"log"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

// Conecta ao RabbitMQ e inicializa filas
func RabbitMQ() error {
	var err error

	// Conectar ao RabbitMQ
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("Erro ao conectar ao RabbitMQ: %v", err)
		return err
	}

	// Abrir canal
	channel, err = conn.Channel()
	if err != nil {
		log.Printf("Erro ao abrir canal: %v", err)
		return err
	}

	// Declarar filas
	_, err = channel.QueueDeclare("video-sent", false, false, false, false, nil)
	if err != nil {
		log.Printf("Erro ao declarar fila video-sent: %v", err)
		return err
	}

	_, err = channel.QueueDeclare("data-sent", false, false, false, false, nil)
	if err != nil {
		log.Printf("Erro ao declarar fila data-sent: %v", err)
		return err
	}

	return nil
}

// Fecha a conexão e o canal
func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}