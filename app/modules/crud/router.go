package crud

import (
	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/middleware"
)

func Router(router fiber.Router) {
	router.Post("/add", middleware.UseCookie, Add)
	router.Patch("/update", middleware.UseCookie, Update)
	router.Delete("/remove/:id<int16>", middleware.UseCookie, Remove)
}
