package crud

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AddClient(c *fiber.Ctx) error {
	_name := c.FormValue("name")
	_mac := strings.ReplaceAll(c.FormValue("mac"), ":", "-")
	_port := c.FormValue("port")
	port, err := strconv.Atoi(_port)
	if err != nil {
		panic(err)
	}

	return c.Status(201).Render("render/table", fiber.Map{"devices": []any{}})
}

// var d Device
// if err := c.BodyParser(&d); err != nil {
// 	return err
// }
