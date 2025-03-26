package crud

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/repository"
)

func Add(c *fiber.Ctx) error {
	_name := c.FormValue("name")
	_mac := strings.ReplaceAll(c.FormValue("mac"), ":", "-")
	_port := c.FormValue("port")
	port, err := strconv.Atoi(_port)
	if err != nil {
		panic(err)
	}

	repository.Write(repository.Device{Name: _name, Mac: _mac, Port: uint16(port)})
	return c.Status(201).Render("render/table", fiber.Map{"devices": repository.DeviceStorage})
}

// var d Device
// if err := c.BodyParser(&d); err != nil {
// 	return err
// }
