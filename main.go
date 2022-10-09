package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppxb/go-fiber/pkg/log"
)

func main() {
	app := fiber.New()

	log.WithContext(nil).Info("server is running...")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})

	err := app.Listen(":5001")
	if err != nil {
		return
	}
}
