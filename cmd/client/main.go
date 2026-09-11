package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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
		log.Fatalf("failed")
	}
	defer conn.Close()

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatalf("failed to welcome client")
	}

	_, _, err = pubsub.DeclareAndBind(conn, routing.ExchangePerilDirect, "pause"+"."+username, routing.PauseKey, pubsub.TransientQueue)
	if err != nil {
		log.Fatalf("failed to declare and bind queue")
	}
	fmt.Printf("Welcome, %s!\n", username)

	gameState := gamelogic.NewGameState(username)

	for {
		input := gamelogic.GetInput()
		if len(input) == 0 {
			fmt.Println("No input detected. Please enter a command.")
			continue
		}
		command := input[0]
		switch command {
			case "spawn":
				

		
	}


	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
