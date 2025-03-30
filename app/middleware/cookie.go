package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/constants"
)

func UseCookie(c *fiber.Ctx) error {
	authCookie := c.Cookies(constants.CookieName)
	if authCookie != constants.AppEnv.CookieSecret {
		headerContent := c.GetReqHeaders()["Hx-Request"]
		if len(headerContent) == 0 {
			return c.Redirect("/login")
		} else {
			c.Set("HX-Redirect", "/login")
			return c.SendStatus(204)
		}
	}
	return c.Next()
}

func MakeCookie() *fiber.Cookie {
	cookie := new(fiber.Cookie)
	cookie.Name = constants.CookieName
	cookie.Value = constants.AppEnv.CookieSecret
	cookie.HTTPOnly = true
	cookie.MaxAge = constants.AppEnv.CookieTTL
	
	return cookie
}
