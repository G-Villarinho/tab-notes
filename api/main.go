package main

import (
	"context"
	"log"
	"time"

	"github.com/g-villarinho/tab-notes-api/app"
	"github.com/g-villarinho/tab-notes-api/clients"
	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/middlewares"
	"github.com/g-villarinho/tab-notes-api/routes"
	"github.com/g-villarinho/tab-notes-api/storages"
)

func main() {
	if err := configs.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := storages.InitDB(ctx)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	queueClient, err := clients.NewRabbitMQPublisher()
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ publisher: %v", err)
	}
	defer queueClient.Close()

	app := app.NewApp(configs.Env.APIPort)

	app.Use(middlewares.CORS)
	app.Use(middlewares.Logging)
	app.Use(middlewares.Recovery)
	app.Use(middlewares.BodySizeLimit)

	router := routes.SetupRoutes(db, queueClient)

	app.RegisterRoutes(router)

	app.Start()
}
