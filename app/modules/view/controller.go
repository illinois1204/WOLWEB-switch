package view

import (
	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/middleware"
	"github.com/illinois1204/WOLWEB-switch/app/repository"
)

func Index(c *fiber.Ctx) error {
	c.Cookie(middleware.MakeCookie())
	return c.Render("index", fiber.Map{"devices": repository.DeviceStorage.ToArray()}, "layouts/main")
}
