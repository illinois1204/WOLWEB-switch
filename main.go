package main

import (
	"github.com/illinois1204/WOLWEB-switch/app/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.RunAppInitialization()
	cmd.RunHttpServer()
}
