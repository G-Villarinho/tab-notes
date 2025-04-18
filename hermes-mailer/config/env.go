package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hermes-mailer/models"
	"github.com/joho/godotenv"
)

var Env models.Environment

func LoadEnv() error {
	if err := godotenv.Load(".env.local"); err != nil {
		return fmt.Errorf("load env: %w", err)
	}

	Env = models.Environment{
		RabbitMQ: models.RabbitMQ{
			URL: os.Getenv("AMQP_URL"),
		},
		SMTP: models.SMTP{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     parseInt(os.Getenv("SMTP_PORT")),
			Username: os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASS"),
		},
	}

	return nil
}

func parseInt(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}
