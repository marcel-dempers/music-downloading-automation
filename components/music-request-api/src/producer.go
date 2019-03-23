package main

import (
  "fmt"
  "log"
  "github.com/streadway/amqp"
  "app/models"
	"strconv"
	"encoding/json"
)

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func produceMessage(message models.Message , config models.RabbitMq) {
	fmt.Println("producer_init started!")
	conn, err := amqp.Dial("amqp://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port) +"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"music-request", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	
	messageString, _ := json.Marshal(message)

 
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(messageString),
		})
	failOnError(err, "Failed to publish a message")

	fmt.Println("producer_init success!")
}


