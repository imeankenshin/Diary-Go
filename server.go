package main

import (
	"first_go/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func main() {
	app := echo.New()
	// middle ware
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	// routing
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello wolrd!")
	})
	// new user
	app.POST("/auth/new", func(c echo.Context) error {
		req := c.Request().Header
		// get user's name, mail, password from request head
		name := req.Get("Name")
		mail := req.Get("Mail")
		password := req.Get("Password")
		// check if name, mail, password are not empty
		if name == "" || mail == "" || password == "" {
			return echo.NewHTTPError(http.StatusBadRequest, ErrorMessage{Message: "name, mail, password can not be empty"})
		}
		data, err := auth.CreateUser(name, mail, password)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, ErrorMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, data.InsertedID)
	})
	// login

	app.POST("/auth", func(c echo.Context) error {
		req := c.Request().Header
		// get user's mail, password from request head
		mail := req.Get("Mail")
		password := req.Get("Password")
		// check if mail, password are not empty
		if mail == "" || password == "" {
			return echo.NewHTTPError(http.StatusBadRequest, ErrorMessage{Message: "mail and password can not be empty"})
		}
		data, err := auth.Login(mail, password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, data)
	})
	app.Static("/assets", "assets")
	// Start server on http://localhost:1000
	app.Logger.Fatal(app.Start(":1000"))
}
