package person_handler

import "github.com/gofiber/fiber/v2"

func GetPersons(c *fiber.Ctx) error {
	return c.SendString("ok")
}
