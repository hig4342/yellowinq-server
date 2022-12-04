package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./client/build")

	app.Static("*", "./client/build/index.html")

	app.Listen(":3000")
}