package server

import (
	"context"
	"github.com/Ovsienko023/reporter/app/core"
	transportHttp "github.com/Ovsienko023/reporter/app/transport/http"
	"github.com/Ovsienko023/reporter/infrastructure"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	httpServer *http.Server

	recordCore *core.Core
}

func New(container *infrastructure.Infrastructure) *App {
	return &App{
		recordCore: core.New(container),
	}
}

func (a *App) Run(apiConfig *configuration.Api) error {
	transport := transportHttp.NewTransport(*a.recordCore, apiConfig)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(transport.StaticHandler)

	r := transportHttp.RegisterHTTPEndpoints(router, transport, apiConfig)

	a.httpServer = &http.Server{
		Addr:           apiConfig.Host + ":" + apiConfig.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
