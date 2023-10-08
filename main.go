package main

import (
	"Rest-api-golang/config"
	"Rest-api-golang/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()

	config.Init()

	routes.BooksRoutes(e)
	routes.UsersRouutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
