package main

import "github.com/labstack/echo/v4"

func main() {
	app := echo.New()
	
	app.Static("/assets", "assets")
	// Start server on http://localhost:1000
	app.Logger.Fatal(app.Start(":100"))
}
