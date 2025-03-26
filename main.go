package main

import (
	"github.com/illinois1204/WOLWEB-switch/app/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd.RunAppInitialization()
	cmd.RunHttpServer()
}
