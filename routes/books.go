package routes

import (
	"Rest-api-golang/controller"

	"github.com/labstack/echo/v4"
)

func BooksRoutes(e *echo.Echo) {

	e.GET("/books", controller.GetBooksController())
	e.GET("/books/:id", controller.GetBookController())
	e.POST("/books", controller.CreateBookControler())
	e.PUT("/books/:id", controller.UpdateBookController())
	e.DELETE("/books/:id", controller.DeleteBookController())
}
