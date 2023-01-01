package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/puskipus/golang-react/bootstrap"
)

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
