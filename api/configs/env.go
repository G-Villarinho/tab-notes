package configs

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/joho/godotenv"
)

var Env models.Environment

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("load env: %w", err)
	}

	Env = models.Environment{
		Env:     getEnv("ENV", "development"),
		APIPort: getEnv("API_PORT", "8080"),
		DB: models.Mysql{
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "3306"),
			DBUser:     getEnv("DB_USER", "root"),
			DBPassword: getEnv("DB_PASSWORD", "root"),
			DBName:     getEnv("DB_NAME", "tabnotes"),
		},
		RedirectURL:    getEnv("REDIRECT_URL", "http://localhost:5173/"),
		APIURL:         getEnv("API_URL", "http://localhost:8080"),
		AllowedOrigins: parseList(getEnv("ALLOWED_ORIGINS", "*")),
		MaxBodySize:    parseInt64(getEnv("MAX_BODY_SIZE", "1048576")), // 1MB
		AMQPURL:        getEnv("AMQP_URL", "amqp://guest:guest@localhost:5672/"),
		Hermes: models.Hermes{
			APIURL: getEnv("HERMES_API_URL", "http://localhost:8888"),
			APIKey: getEnv("HERMES_API_KEY", ""),
		},
	}

	privateKey, err := loadKeyFromFile(os.Getenv("KEY_ECDSA_PRIVATE"))
	if err != nil {
		return fmt.Errorf("load private key: %w", err)
	}

	publicKey, err := loadKeyFromFile(os.Getenv("KEY_ECDSA_PUBLIC"))
	if err != nil {
		return fmt.Errorf("load public key: %w", err)
	}

	Env.Key = models.Key{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func parseList(val string) []string {
	parts := strings.Split(val, ",")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}

func parseInt64(val string) int64 {
	n, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

func loadKeyFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return strings.TrimSpace(string(data)), nil
}
