package repository

import (
	"balbibe/database"
	"balbibe/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllCourse(course []models.Course, c echo.Context) error {
	err := database.DB.Find(&course).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": 200,
		"data":    &course,
	})
}

func FindCourseByID(course []models.Course, id string, c echo.Context) error {
	database.DB.Model(&course).Where("id = ?", id).Update("view", gorm.Expr("view + ?", 1))
	err := database.DB.Where("id = ?", id).Find(&course).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": 200,
		"data":    &course,
	})
}

func AddCourse(course models.Course, c echo.Context) error {
	err := database.DB.Create(&course).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": 200,
		"data":    &course,
	})
}

func GetHistoryTransaction(history []models.HistoryTransactionCourse, id string, c echo.Context) error {
	err := database.DB.Raw("SELECT t.id, u.nama, t.book_title,t.buy_at, b.file FROM user u INNER JOIN transaction_book t ON u.id = t.id_user INNER JOIN book b ON b.id = t.id_book WHERE u.id = ?", id).Scan(&history).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": 200,
		"data":    &history,
	})
}

func EditHistoryTransaction(history models.HistoryTransactionCourse, modelCourse models.TransactionBook, id string, idhistory string, c echo.Context) error {
	err := database.DB.Model(&modelCourse).Where("id = ?", idhistory).Updates(&modelCourse).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Update",
	})
}
