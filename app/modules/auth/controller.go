package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/constants"
	"github.com/illinois1204/WOLWEB-switch/app/middleware"
)

func Login(c *fiber.Ctx) error {
	return c.Render("auth", fiber.Map{}, "layouts/main")
}

func SignIn(c *fiber.Ctx) error {
	pass := c.FormValue("password")
	pass = strings.TrimSpace(pass)

	if pass == "" || pass != constants.AppEnv.Password {
		return c.SendStatus(401)
	}

	c.Cookie(middleware.MakeCookie())
	return c.Redirect("/", fiber.StatusMovedPermanently)
}

func SignOut(c *fiber.Ctx) error {
	c.ClearCookie(constants.CookieName)
	c.Set("HX-Redirect", "/login")
	return c.SendStatus(204)
}
