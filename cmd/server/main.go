package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")

	const amqpURI = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}

	fmt.Println("Connection successful")
	gamelogic.PrintServerHelp()

outerLoop:
	for {
		arr := gamelogic.GetInput()
		if len(arr) == 0 {
			fmt.Println("No input detected. Please enter a command.")
			continue
		}
		firstWord := arr[0]

		switch firstWord {
		case "pause":
			fmt.Println("Sending a pause message")
			err := pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
				IsPaused: true,
			})
			if err != nil {
				log.Fatalf("failed to send pause message: %v", err)
			}

		case "resume":
			fmt.Println("Sending a resume message")
			err := pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
				IsPaused: false,
			})
			if err != nil {
				log.Fatalf("failed to send resume message: %v", err)
			}

		case "quit":
			fmt.Println("Quitting the server")
			break outerLoop

		default:
			fmt.Println("Can't understand command")
		}

		fmt.Println("Shutting down")

	}
}
