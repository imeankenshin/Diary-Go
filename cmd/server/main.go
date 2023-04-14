package server

import (
	"first_go/internal/auth"
	"first_go/internal/diary"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func App() *echo.Echo {
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
	app.POST("/auth/new", auth.CreateUser)
	app.POST("/auth", auth.Login)
	// 認可しないとアクセスできないルート
	userOnly := app.Group("/diary")
	userOnly.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("nothing"),
	}))
	userOnly.GET("", diary.GetDiary)
	userOnly.POST("", diary.PostDiary)
	return app
}
