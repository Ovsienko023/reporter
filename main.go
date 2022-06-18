package main

import (
	"github.com/Ovsienko023/reporter/server"
	"log"
)

func main() {
	//application := app.InitApp()

	//http.ListenAndServe(":8888", application)
	app := server.NewApp()
	if err := app.Run("8888"); err != nil { // todo add config
		log.Fatalf("%s", err.Error())
	}
}
