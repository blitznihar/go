package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

type Client struct {
	C               *websocket.Conn
	Id              string
	SecWebsocketKey string
}
type Message struct {
	Id            string
	ClientId      string
	FirstName     string
	LastName      string
	CreatedTime   time.Time
	ContactNumber string
}

var clients = make(map[string]Client)

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {

		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		websocketkey := c.Headers("Sec-Websocket-Key")
		clients[c.Params("id")] = Client{
			C:               c,
			Id:              c.Params("id"),
			SecWebsocketKey: websocketkey,
		}
		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("from: %s recv: %s ", websocketkey, msg)
			//TODO: Send Message with ClientId to Queue

			//client := clients[c.Params("id")]
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
			//forever := make(chan bool)
			go func() {
				for msg := range msgs {
					var b Message
					err := json.Unmarshal([]byte(msg.Body), &b)
					if err != nil {
						log.Panic(err)
					}
					response := []byte("messageId" + fmt.Sprint(1) + " from: " + websocketkey + " recv: " + string(msg.Body))
					client, ok := clients[b.ClientId]
					if ok {
						if err = client.C.WriteMessage(mt, response); err != nil {
							log.Println("write:", err)
							break
						}
					}
					fmt.Printf("Recieved Message: %s\n", msg.Body)

				}
			}()

			fmt.Println("Successfully Connected to our RabbitMQ Instance")
			fmt.Println(" [*] - Waiting for messages")
			//<-forever

			defer ch.Close()
			defer conn.Close()

		}

	}))

	log.Fatal(app.Listen(":3001"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
