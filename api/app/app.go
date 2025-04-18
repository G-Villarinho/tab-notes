package app

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	router      http.Handler
	Port        string
	middlewares []func(http.Handler) http.Handler
}

func NewApp(port string) *App {
	return &App{
		Port:        port,
		middlewares: []func(http.Handler) http.Handler{},
	}
}

func (a *App) Use(mw ...func(http.Handler) http.Handler) {
	a.middlewares = append(a.middlewares, mw...)
}

func (a *App) RegisterRoutes(router http.Handler) {
	a.router = router
}

func (a *App) Start() {
	if a.router == nil {
		log.Fatal("router not registered: call RegisterRoutes() before Start()")
	}

	handler := a.router
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		handler = a.middlewares[i](handler)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.Port),
		Handler: handler,
	}

	log.Printf("🔥 Starting server on port %s\n", a.Port)
	log.Fatal(server.ListenAndServe())
}
