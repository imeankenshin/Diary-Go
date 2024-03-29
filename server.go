package main

import (
	"first_go/cmd/server"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func main() {
	app := server.App()
	// Start server on https://localhost:1000
	app.Logger.Fatal(app.StartTLS(":1000", "certificate/cert.pem", "certificate/key.pem"))
}
