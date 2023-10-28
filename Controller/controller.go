package Controller

import (
	"awesomeProject/Controller/User"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) error {
	userSetupErr := User.SetupHandlers(app)
	if userSetupErr != nil {
		return userSetupErr
	}

	return nil
}
