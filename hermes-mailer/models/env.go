package models

type Environment struct {
	RabbitMQ
	SMTP
}

type RabbitMQ struct {
	URL string
}

type SMTP struct {
	Host     string
	Port     int
	Username string
	Password string
}
