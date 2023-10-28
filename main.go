package main

import (
	"awesomeProject/Controller/User"
	json2 "encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
)

var fakeDB = []User.SUser{
	{
		Id:   1,
		Age:  20,
		Name: "Pedro",
		Wallet: User.SWallet{
			Money: 0,
			Cards: []User.SCard{
				{
					Bank:   "Inter",
					Secret: 123,
				},
				{
					Bank:   "C6",
					Secret: 321,
				},
			},
		},
	},
}

func FindById(id int, db []User.SUser) *User.SUser {
	for _, user := range db {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

func GetById(context *fiber.Ctx) error {
	id, errParams := context.ParamsInt("id")
	if errParams != nil {
		context.Status(fiber.StatusBadRequest)
		return errParams
	}

	userRaw := FindById(id, fakeDB)
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
	json, parseErr := json2.Marshal(fakeDB)
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

func main() {
	println("Hello, Trainees!")
	app := fiber.New()

	app.Get("api/user/all", GetAll)
	app.Get("api/user/:id", GetById)

	app.Listen(":8000")
}
