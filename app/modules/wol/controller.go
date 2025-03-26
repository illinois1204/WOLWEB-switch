package wol

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func WakeUp(c *fiber.Ctx) error {
	// return c.Render("index", fiber.Map{"devices": repository.DeviceStorage}, "layouts/main")
	fmt.Println(c.Params("id"))
	return c.SendStatus(200)
}
