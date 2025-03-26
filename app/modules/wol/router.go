package wol

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router) {
	router.Put("/wake-up/:id<int16>", WakeUp)
}
