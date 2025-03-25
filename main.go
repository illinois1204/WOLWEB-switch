package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
	"github.com/illinois1204/WOLWEB-switch/app/constants"
	"github.com/illinois1204/WOLWEB-switch/app/modules/crud"
	"github.com/illinois1204/WOLWEB-switch/app/modules/view"
	"github.com/illinois1204/WOLWEB-switch/app/service"
	"github.com/joho/godotenv"
)

// type Device struct {
// 	Id   uint
// 	Name string `form:"name"`
// 	Mac  string `form:"mac"`
// 	Port uint16 `form:"port"`
// }
// var rows []Device
// rows = append(rows, Device{Id: 1, Name: "NAS", Mac: "28-C6-8E-36-DC-38", Port: 9})
// rows = append(rows, Device{Id: 2, Name: "Laptop", Mac: "29-C6-8E-36-DC-38", Port: 9})
// rows = append(rows, Device{Id: 3, Name: "PC", Mac: "28-C6-FE-36-DC-38", Port: 7})
// rows = append(rows, Device{Id: 4, Name: "Server", Mac: "28-C6-8E-74-DC-38", Port: 9})
// rows = append(rows, Device{Id: 5, Name: "Runner", Mac: "18-1D-14-70-A0-21", Port: 9})

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	constants.Load()
	_, _, lastWrittenFile := service.ListStoreFiles()
	lastIndex := 0
	if lastWrittenFile != "" {
		lastIndex = service.GetLastFileIndex(lastWrittenFile)
	}
	service.MakeCounter(int64(lastIndex))

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
