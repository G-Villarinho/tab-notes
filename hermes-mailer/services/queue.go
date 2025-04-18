package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hermes-mailer/models"
)

type QueueService interface {
	Start() error
}

type queueService struct {
	queueClient     models.QueueConsumer
	emailService    EmailService
	processingQueue string
}

func NewQueueService(
	client models.QueueConsumer,
	emailService EmailService,
	queueName string) QueueService {
	return &queueService{
		queueClient:     client,
		emailService:    emailService,
		processingQueue: queueName,
	}
}

func (q *queueService) Start() error {
	msgs, err := q.queueClient.Consume(q.processingQueue)
	if err != nil {
		return err
	}

	for msg := range msgs {
		go q.handleMessage(msg.Body)
	}

	return nil
}

func (q *queueService) handleMessage(body []byte) {
	var email models.Email
	if err := json.Unmarshal(body, &email); err != nil {
		log.Println("❌ Erro ao deserializar:", err)
		return
	}

	if err := q.emailService.Send(context.Background(), email); err != nil {
		log.Println("❌ Erro ao enviar e-mail:", err)
		return
	}

	log.Println("✅ Email enviado com sucesso")
}
