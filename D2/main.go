package main

import (
	"agmc_d2/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting .env file")
	}

	e := echo.New()
	routes.SetRouter(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
