package crud

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router) {
	router.Post("/add", AddClient)
}
