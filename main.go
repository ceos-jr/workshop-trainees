package main

import (
	"awesomeProject/Controller"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	println("Hello, Trainees!")
	app := fiber.New()

	controllerSetupErr := Controller.Setup(app)
	if controllerSetupErr != nil {
		log.Fatalf(controllerSetupErr.Error())
	}

	appErr := app.Listen(":8000")
	if appErr != nil {
		log.Fatalf(appErr.Error())
	}
}
