package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type Message struct {
	Id            string
	FirstName     string
	LastName      string
	CreatedTime   time.Time
	ContactNumber string
}

func main() {

	// Handle any errors if we were unable to create the queue
	go sendMessage()

	receiveMessage()

}

func receiveMessage() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(" [*] - Waiting for messages - Started")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)

		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever

	defer ch.Close()
	defer conn.Close()
}

func sendMessage() {

	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		time.Sleep(1 * time.Second)
		fmt.Println(err)
		panic(1)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 100; i++ {
		message := Message{
			Id:            uuid.New().String(),
			FirstName:     "John",
			LastName:      "Doe",
			CreatedTime:   time.Now(),
			ContactNumber: "123456789",
		}
		log.Println(message)
		b, err := json.Marshal(message)
		if err != nil {
			log.Panic(err)
		}

		err = ch.Publish(
			"",
			"TestQueue",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(b),
			},
		)

		if err != nil {
			fmt.Println(err)
		}

	}
	defer ch.Close()
	defer conn.Close()

}
