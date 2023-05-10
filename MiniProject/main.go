package main

import (
	"MiniProject/database"
	"MiniProject/routes"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "8080"

func main() {
	database.InitDB()

	app := echo.New()
	routes.Routes(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	appPort := fmt.Sprintf(":%s", port)

	app.Logger.Fatal(app.Start(appPort))

}
