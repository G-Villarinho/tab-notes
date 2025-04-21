package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/hermes-mailer/clients"
	"github.com/hermes-mailer/config"
	"github.com/hermes-mailer/handlers"
	"github.com/hermes-mailer/middlewares"
	"github.com/hermes-mailer/services"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("load enviroments: %v", err)
	}

	rabbit, err := clients.NewRabbitMQClient(config.Env.RabbitMQ.URL)
	if err != nil {
		log.Fatal("connect RabbitMQ:", err)
	}
	defer rabbit.Close()

	smtpClient := clients.NewSMTPEmailSenderClient()

	queueService := services.NewQueueService(rabbit, config.Env.API.QueueName)
	emailService := services.NewEmailService(smtpClient, queueService)

	emailHandler := handlers.NewEmailHandler(emailService)

	mux := http.NewServeMux()
	mux.Handle("/send-email", middlewares.RequireAPIKey()(http.HandlerFunc(emailHandler.SendEmail)))

	port := strconv.Itoa(config.Env.API.Port)
	log.Printf("🔥 Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
