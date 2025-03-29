package crud

import (
	"fmt"
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
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	object := repository.Device{Name: _name, Mac: _mac, Port: uint16(port)}
	index, err := repository.Write(object)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	repository.DeviceStorage.Add(index, object)
	return c.Status(201).Render("render/table", fiber.Map{"devices": repository.DeviceStorage.ToArray()})
}

func Update(c *fiber.Ctx) error {
	var d repository.Device
	if err := c.BodyParser(&d); err != nil {
		return c.Status(500).SendString("Oops, something went wrong")
	}

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	if err = repository.DeviceStorage.Update(uint(id), d); err != nil {
		return c.Status(500).SendString("Oops, something went wrong")
	}

	return c.Status(200).Render("render/table", fiber.Map{"devices": repository.DeviceStorage.ToArray()})
}

func Remove(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	repository.DeviceStorage.Remove(uint(id))
	return c.Status(200).Render("render/table", fiber.Map{"devices": repository.DeviceStorage.ToArray()})
}
