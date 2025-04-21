package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hermes-mailer/clients"
	"github.com/hermes-mailer/config"
	"github.com/hermes-mailer/models"
	"github.com/hermes-mailer/services"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("loading .env: %v", err)
	}

	rabbit, err := clients.NewRabbitMQClient(config.Env.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}
	defer rabbit.Close()

	smtpClient := clients.NewSMTPEmailSenderClient()

	queueService := services.NewQueueService(rabbit, config.Env.API.QueueName)
	emailService := services.NewEmailService(smtpClient, queueService)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go waitForShutdown(cancel)

	log.Println("📡 Worker is running and listening for emails...")

	msgs, err := queueService.Consume(ctx, config.Env.API.QueueName)
	if err != nil {
		log.Fatalf("consume queue: %v", err)
	}

	for body := range msgs {
		go func(b []byte) {
			var email models.Email
			if err := json.Unmarshal(b, &email); err != nil {
				log.Printf("unmarshal email: %v", err)
				return
			}

			if err := emailService.SendEmail(ctx, email); err != nil {
				log.Printf("send email: %v", err)
				return
			}

			log.Printf("✅ Email sent")
		}(body)
	}
}

func waitForShutdown(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	log.Println("🛑 Shutting down worker...")
	cancel()
}
