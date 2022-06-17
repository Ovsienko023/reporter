package main

import (
	"github.com/Ovsienko023/reporter/app"
	"net/http"
)

func main() {
	application := app.InitApp()

	http.ListenAndServe(":8888", application)
}
