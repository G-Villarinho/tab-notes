package main

import (
	"log"

	"github.com/hermes-mailer/clients"
	"github.com/hermes-mailer/config"
)

func main() {
	config.LoadEnv()

	rabbit, err := clients.NewRabbitMQClient(config.Env.RabbitMQ.URL)
	if err != nil {
		log.Fatal("❌ Erro ao conectar no RabbitMQ:", err)
	}
	defer rabbit.Close()

}
