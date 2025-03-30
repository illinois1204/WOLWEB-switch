package auth

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router) {
	router.Get("/login", Login)
	router.Post("/sign-in", SignIn)
	router.Delete("/sign-out", SignOut)
}
