package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/illinois1204/WOLWEB-switch/app/constants"
	"github.com/illinois1204/WOLWEB-switch/app/modules/crud"
	"github.com/illinois1204/WOLWEB-switch/app/modules/view"
)

func RunHttpServer() {
	port := constants.AppEnv.Port
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		AppName:       "WOL web switcher",
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		Views:         engine,
	})

	app.Static("/static", "./public")
	view.Router(app)
	crud.Router(app.Group("/manage"))

	fmt.Println("Application started!")
	app.Listen(fmt.Sprintf("0.0.0.0:%s", port))
}
