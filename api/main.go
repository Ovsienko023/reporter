package main

import (
	"flag"
	"fmt"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"log"

	"github.com/Ovsienko023/reporter/server"
)

func main() {
	flag.Parse()

	cfg, err := configuration.NewConfig()

	if err != nil {
		log.Fatalf("Could not read configuration file with error: %+v", err)
	}

	fmt.Printf("Running on: %s:%s \n", cfg.Api.Host, cfg.Api.Port)

	app := server.NewApp(cfg)
	if err := app.Run(&cfg.Api); err != nil {
		log.Fatalf("%s", err.Error())
	}
}