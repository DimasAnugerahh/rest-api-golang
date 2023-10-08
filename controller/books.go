package controller

import (
	"Rest-api-golang/config"
	"Rest-api-golang/model"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// create user
func CreateBookControler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books := model.Book{}
		c.Bind(&books)

		if err := config.DB.Save(&books).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new book",
			"users":   books,
		})
	}
}

// get books
func GetBooksController() echo.HandlerFunc {
	return func(c echo.Context) error {
		books := []model.Book{}

		if err := config.DB.Find(&books).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success gel all book",
			"users":   books,
		})
	}
}

// get book by id
func GetBookController() echo.HandlerFunc {

	return func(c echo.Context) error {
		books := model.Book{}
		var idParam = c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		if err := config.DB.Where("id=?", id).Find(&books).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"users":   books,
		})
	}
}

// update book by id
func UpdateBookController() echo.HandlerFunc {
	return func(c echo.Context) error {
		// your solution here
		books := model.Book{}
		var idParam = c.Param("id")
		c.Bind(&books)
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "invalid user input",
			})
		}

		var qry = config.DB.Table("books").Where("id=?", id).Updates(map[string]any{
			"judul":    books.Judul,
			"penulis":  books.Penulis,
			"penerbit": books.Penerbit,
		})
		if dataCount := qry.RowsAffected; dataCount < 1 {
			logrus.Error("Model : Update error, ", "no data affected")
		}
		if err := qry.Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())

		}
		return c.JSON(http.StatusCreated, map[string]any{
			"messege": "success",
			"data":    books,
		})
	}
}

// delete books
func DeleteBookController() echo.HandlerFunc {
	return func(c echo.Context) error {
		books := model.Book{}

		c.Bind(&books)
		var idParam = c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		if err := config.DB.Where("id=? and deleted_at is null", id).Delete(&books).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"message": "Deleted",
			"books":   books,
		})
	}
}
