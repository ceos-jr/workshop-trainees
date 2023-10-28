package User

import (
	"github.com/gofiber/fiber/v2"
)

type (
	SCard struct {
		Bank   string `json:"bank"`
		Secret int    `json:"secret"`
	}

	SWallet struct {
		Money int     `json:"money"`
		Cards []SCard `json:"cards"`
	}

	SUser struct {
		Id     int     `json:"id"`
		Age    int     `json:"age"`
		Name   string  `json:"name"`
		Wallet SWallet `json:"wallet"`
	}

	BaseHandler struct {
		CRUD REST
	}

	REST interface {
		Post(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		GetById(c *fiber.Ctx) error
		Put(c *fiber.Ctx) error
		Delete(c *fiber.Ctx) error
	}
)
