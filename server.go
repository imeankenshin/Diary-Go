package main

import (
	"first_go/routes"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func main() {
	app := routes.App()
	// Start server on https://localhost:1000
	app.Logger.Fatal(app.Start(":1000"))
}
