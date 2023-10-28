package User

import (
	"awesomeProject/Model"
	"awesomeProject/Model/User"
	json2 "encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func findById(id int, db []User.SUser) *User.SUser {
	for _, user := range db {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

func Post(context *fiber.Ctx) error {
	return nil
}

func GetById(context *fiber.Ctx) error {
	id, errParams := context.ParamsInt("id")
	if errParams != nil {
		context.Status(fiber.StatusBadRequest)
		return errParams
	}

	userRaw := findById(id, Model.FakeDB)
	if userRaw == nil {
		context.Status(fiber.StatusNotFound)
		return errors.New("not found")
	}

	user := *userRaw
	userBytes, parseErr := json2.Marshal(user)

	if parseErr != nil {
		context.Status(fiber.StatusBadRequest)
		return parseErr
	}

	sendErr := context.Send(userBytes)
	if sendErr != nil {
		return sendErr
	}

	return nil
}

func GetAll(context *fiber.Ctx) error {
	json, parseErr := json2.Marshal(Model.FakeDB)
	if parseErr != nil {
		context.Status(fiber.StatusBadRequest)
		return parseErr
	}

	sendErr := context.Send(json)
	if sendErr != nil {
		return sendErr
	}
	return nil
}
func Put(context *fiber.Ctx) error {
	return nil
}

func Delete(context *fiber.Ctx) error {
	return nil
}

func SetupHandlers(app *fiber.App) error {
	app.Get("api/user/all", GetAll)
	app.Get("api/user/:id", GetById)

	return nil
}
