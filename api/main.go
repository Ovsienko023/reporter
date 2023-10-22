package main

import (
	"flag"
	"fmt"
	"github.com/Ovsienko023/reporter/infrastructure"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/Ovsienko023/reporter/server"
	"log"
)

// @title Reporter API
// @version 0.0.1
// @description This is a report server.

// @BasePath /api/v1
func main() {
	flag.Parse()

	cfg, err := configuration.NewConfig()
	//db, err := postgres.New(&cfg.Db)
	//if err != nil {
	//	panic(err)
	//}

	if err != nil {
		log.Fatalf("Could not read configuration file with error: %+v", err)
	}

	fmt.Printf("Running on: %s:%s \n", cfg.Api.Host, cfg.Api.Port)

	inf, err := infrastructure.New(*cfg)
	if err != nil {
		panic(err)
	}

	app := server.New(inf)
	if err := app.Run(&cfg.Api); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
