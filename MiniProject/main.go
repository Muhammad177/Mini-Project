package main

import (
	"MiniProject/database"
	"MiniProject/routes"
)

func main() {

	database.InitDB()
	database.InitialMigration()
	e := routes.New()

	e.Start("8000")

}
