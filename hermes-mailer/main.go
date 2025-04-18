package main

import (
	"log"
	"os"

	"github.com/hermes-mailer/clients"
	"github.com/hermes-mailer/config"
	"github.com/hermes-mailer/models"
	"github.com/hermes-mailer/services"
)

func main() {
	config.LoadEnv()

	rabbit, err := clients.NewRabbitMQClient(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatal("❌ Erro ao conectar no RabbitMQ:", err)
	}
	defer rabbit.Close()

	emailClient := clients.NewSMTPEmailSenderClient()
	emailService := services.NewEmailService(emailClient)

	var queueConsumer = models.QueueConsumer(rabbit)
	queueService := services.NewQueueService(queueConsumer, emailService, "email_queue")

	log.Println("📬 Hermes Mailer Service iniciado com sucesso.")
	log.Println("📡 Aguardando mensagens na fila: email_queue...")

	if err := queueService.Start(); err != nil {
		log.Fatal("❌ Erro ao iniciar QueueService:", err)
	}
}
