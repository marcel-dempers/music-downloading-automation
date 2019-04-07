package main


import (
	"fmt"
	amqp "github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
	"app/models"
	"encoding/json"
	 
)

var appConfig *models.Configuration 

func failOnErrorRetry(err error, msg string) {
	if err != nil {
	  fmt.Println("Error occurred: " + msg)
	}
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}
 
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

func main() {
	c := GetConfiguration()
	appConfig = c 
	//var conn *amqp.Connection
	//var ch *amqp.Channel
	//var q *amqp.Queue

	queueCfg := appConfig.RabbitMq
	duration, _ := time.ParseDuration("10s")
	retry(3, duration, func() (err error){
		conn, err := amqp.Dial("amqp://" + queueCfg.Username + ":" + queueCfg.Password + "@" + queueCfg.Host + ":" + strconv.Itoa(queueCfg.Port) +"/")
		failOnErrorRetry(err, "Failed to connect to RabbitMQ")
		
		if err == nil {
			ch, err := conn.Channel()
			failOnErrorRetry(err, "Failed to open a channel")
	
			q, err := ch.QueueDeclare(
				"music-request", // name
				false,   // durable
				false,   // delete when unused
				false,   // exclusive
				false,   // no-wait
				nil,     // arguments
			)
			failOnErrorRetry(err, "Failed to declare a queue")
	
			fmt.Println("Channel and Queue established")
			fmt.Println(q)
			
			defer conn.Close()
			defer ch.Close()

			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,   // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			  )
			  failOnErrorRetry(err, "Failed to register a consumer")
		
			  forever := make(chan bool)
		
			  go func() {
				for d := range msgs {
					log.Printf("Received a message: %s", d.Body)
					message := models.Message{}
					 
					if err := json.Unmarshal(d.Body, &message); err != nil {
						panic(err)
					}
					err = ProcessMessage(message)
					if err != nil {
						panic(err)
					}

					d.Ack(false)
				}
			  }()
			  
			  fmt.Println("Running...")
			  <-forever
		}

		
		
		return err
	})
}
