package routes

import (
	"Rest-api-golang/controller"

	"github.com/labstack/echo/v4"
)

func UsersRouutes(e *echo.Echo) {

	e.GET("/users", controller.GetUsersController())
	e.GET("/users/:id", controller.GetUserController())
	e.POST("/users", controller.CreateUserController())
	e.DELETE("/users/:id", controller.DeleteUserController())
	e.PUT("/users/:id", controller.UpdateUserController())
}
