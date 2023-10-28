package User

import (
	"github.com/gofiber/fiber/v2"
)

type (
	REST interface {
		Post(c *fiber.Ctx) error
		GetAll(c *fiber.Ctx) error
		GetById(c *fiber.Ctx) error
		Put(c *fiber.Ctx) error
		Delete(c *fiber.Ctx) error
	}
)
