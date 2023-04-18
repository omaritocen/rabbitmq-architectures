package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"rabbitmqhello/utils"
	"time"
)

/*
go run emit_direct.go "Hello Info" info
go run emit_direct.go "Hello warning" warning
go run emit_direct.go "Hello error" error
*/
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to Open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := utils.BodyFrom(os.Args)

	err = ch.PublishWithContext(ctx,
		"logs_direct",
		utils.SeverityFrom(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
}
