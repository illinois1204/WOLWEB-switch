package wol

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/repository"
	"github.com/illinois1204/WOLWEB-switch/app/service"
)

func WakeUp(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	device := repository.DeviceStorage[uint(id)]
	if err := service.WakeUp(device.Mac, device.Port); err != nil {
		fmt.Printf("WOL sending error. Detail:\n%v\n", err)
	}
	return c.SendStatus(204)
}
