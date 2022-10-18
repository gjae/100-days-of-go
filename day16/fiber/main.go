package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func handler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"name": c.Params("user"),
	})
}

func main() {
	engine := django.New("views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/:user", handler)

	app.Hooks().OnListen(func (r fiber.Route) error {
		log.Println("Start")
		return nil
	})

	app.Listen(":8000")
}