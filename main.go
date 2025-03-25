package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/joho/godotenv"
)

type Device struct {
	Id   uint
	Name string `form:"name"`
	Mac  string `form:"mac"`
	Port uint16 `form:"port"`
}

func main() {
	var rows []Device
	rows = append(rows, Device{Id: 1, Name: "NAS", Mac: "28-C6-8E-36-DC-38", Port: 9})
	rows = append(rows, Device{Id: 2, Name: "Laptop", Mac: "29-C6-8E-36-DC-38", Port: 9})
	rows = append(rows, Device{Id: 3, Name: "PC", Mac: "28-C6-FE-36-DC-38", Port: 7})
	rows = append(rows, Device{Id: 4, Name: "Server", Mac: "28-C6-8E-74-DC-38", Port: 9})
	rows = append(rows, Device{Id: 5, Name: "Runner", Mac: "18-1D-14-70-A0-21", Port: 9})

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		AppName:       "WOL web switcher",
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		Views:         engine,
	})

	app.Static("/static", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"devices": rows}, "layouts/main")
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		_name := c.FormValue("name")
		_mac := strings.ReplaceAll(c.FormValue("mac"), ":", "-")
		_port := c.FormValue("port")

		port, err := strconv.Atoi(_port)
		if err != nil {
			panic(err)
		}

		rows = append(rows, Device{Id: 6, Name: _name, Mac: _mac, Port: uint16(port)})
		return c.Status(201).Render("render/table", fiber.Map{"devices": rows})

		// var d Device
		// if err := c.BodyParser(&d); err != nil {
		// 	return err
		// }
	})

	fmt.Println("Application started!")
	app.Listen(fmt.Sprintf("0.0.0.0:%s", port))
}
