package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Ovsienko023/reporter/internal/config"
	"github.com/Ovsienko023/reporter/server"
)

func main() {
	flag.Parse()

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Could not read config file with error: %+v", err)
	}

	fmt.Println(*cfg)

	app := server.NewApp()
	if err := app.Run(&cfg.Api); err != nil { // todo add config
		log.Fatalf("%s", err.Error())
	}
}
