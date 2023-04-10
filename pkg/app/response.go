package app

import "github.com/labstack/echo/v4"

type ErrorMessage struct {
	Message string `json:"message"`
}

func ServerError(c echo.Context, err error) error {
	return c.JSON(500, ErrorMessage{Message: err.Error()})
}
