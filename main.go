package main

import (
	""

	"github.com/gofiber/fiber/v2"
	cmd "github.com/tooolbox/esbuild/src/esbuild/main"
)

type url struct {
	prefix string
	root   string
}

var urls []url = []url{
	{prefix: "/", root: "./public/page"},
	{prefix: "/detail", root: "./public/page/"},
}

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})
	for i := 0; i >= len(urls)-1; i++ {
		app.Static(urls[i].prefix, urls[i].root)
	}

	build.setupEsbuild()
	cmd.Run()
	// app.Listen(":4000")
}
