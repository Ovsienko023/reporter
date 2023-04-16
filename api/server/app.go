package server

import (
	"context"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/repository"
	transportHttp "github.com/Ovsienko023/reporter/app/transport/http"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/Ovsienko023/reporter/infrastructure/logger"
	"io"
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

	Core *core.Core
}

func NewApp(cnf *configuration.Config) *App {
	client, _ := repository.New(&cnf.Db)

	lgr, _ := logger.New(nil)

	return &App{
		Core: core.New(
			client,
			&core.Config{Logger: lgr},
		),
	}
}

func (a *App) Run(apiConfig *configuration.Api) error {
	router := chi.NewRouter()
	//lgr := a.Core.GetLogger()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	r := transportHttp.RegisterHTTPEndpoints(router, *a.Core, apiConfig)

	a.httpServer = &http.Server{
		Addr:           apiConfig.Host + ":" + apiConfig.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		//ErrorLog:       logger.New(&fwdToZapWriter{logger}, "", 0),
		//ErrorLog: logger.New(serverJsonWriter{}, "", 0),
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

type serverJsonWriter struct {
	io.Writer
}

//// ListenAndServeTLS - with custom log Writer
//func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
//	server := &http.Server{
//		Addr: addr,
//		Handler: handler,
//
//	}
//}
//
//func (w serverJsonWriter) Write(p []byte) (n int, err error) {
//	//fw.logger.Errorw(string(p))
//	return len(p), nil
//}

//type fwdToZapWriter struct {
//	logger *zap.SugaredLogger
//}

//func (fw *fwdToZapWriter) Write(p []byte) (n int, err error) {
//	fw.logger.Errorw(string(p))
//	return len(p), nil
//}
