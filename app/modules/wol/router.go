package wol

import (
	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/middleware"
)

func Router(router fiber.Router) {
	router.Put("/wake-up/:id<int16>", middleware.UseCookie, WakeUp)
}
