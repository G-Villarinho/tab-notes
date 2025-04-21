package models

type Environment struct {
	API      API
	RabbitMQ RabbitMQ
	SMTP     SMTP
}

type API struct {
	QueueName string
	APIKey    string
	Port      int
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
