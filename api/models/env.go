package models

type Environment struct {
	Env            string
	APIPort        string
	DB             Mysql
	Key            Key
	AllowedOrigins []string
	MaxBodySize    int64
	RedirectURL    string
	APIURL         string
	AMQPURL        string
}

type Mysql struct {
	DBHost     string
	DBName     string
	DBPort     string
	DBUser     string
	DBPassword string
}

type Key struct {
	PrivateKey string
	PublicKey  string
}
