package controller

import (
	"Rest-api-golang/config"
	"Rest-api-golang/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

// get all users
func GetUsersController() echo.HandlerFunc {
	users := []model.User{}
	return func(c echo.Context) error {

		if err := config.DB.Find(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all",
			"users":   users,
		})
	}
}

// get user by id
func GetUserController() echo.HandlerFunc {

	return func(c echo.Context) error {
		users := model.User{}
		var idParam = c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		if err := config.DB.Where("id=?", id).Find(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"users":   users,
		})
	}
}

// create new user
func CreateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := model.User{}
		c.Bind(&users)

		if err := config.DB.Save(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new users",
			"users":   users,
		})
	}
}

// delete user by id
func DeleteUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := model.User{}

		c.Bind(&users)
		var idParam = c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		if err := config.DB.Where("id=? and deleted_at is null", id).Delete(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"message": "Deleted",
			"users":   users,
		})
	}
}

// update user by id
func UpdateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// your solution here
		users := model.User{}
		var idParam = c.Param("id")
		c.Bind(&users)
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "invalid user input",
			})
		}

		var qry = config.DB.Table("users").Where("id=?", id).Updates(map[string]any{
			"name":     users.Name,
			"email":    users.Email,
			"password": users.Email,
		})
		if dataCount := qry.RowsAffected; dataCount < 1 {
			logrus.Error("Model : Update error, ", "no data affected")
		}
		if err := qry.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())

		}
		return c.JSON(http.StatusCreated, map[string]any{
			"messege": "success",
			"data":    users,
		})
	}
}
