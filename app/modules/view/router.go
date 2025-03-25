package view

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router) {
	router.Get("/", Index)
}
