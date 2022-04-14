package handler

import (
	"balbibe/models"
	"balbibe/repository"

	"github.com/labstack/echo/v4"
)

func GetCourseAll(c echo.Context) error {
	var course []models.Course

	return repository.GetAllCourse(course, c)
}

func GetCourseById(c echo.Context) error {
	id := c.Param("id")
	var course []models.Course

	return repository.FindCourseByID(course, id, c)
}

func AddReviewById(c echo.Context) error {
	review := models.ReviewParenting{}
	if err := c.Bind(&review); err != nil {
		return err
	}
	return repository.AddReviewById(review, c)
}

func CreateCourse(c echo.Context) error {
	var courseModel models.Course

	if err := c.Bind(&courseModel); err != nil {
		return err
	}

	return repository.AddCourse(courseModel, c)
}

func GetHistoryTransaction(c echo.Context) error {
	var history []models.HistoryTransactionCourse
	id := c.Param("id")

	return repository.GetHistoryTransaction(history, id, c)
}

func EditHistoryTransaction(c echo.Context) error {
	var history models.HistoryTransactionCourse
	var modelCourse models.TransactionBook
	id := c.Param("id")
	idhistory := c.Param("idhistory")
	if err := c.Bind(&modelCourse); err != nil {
		return err
	}
	return repository.EditHistoryTransaction(history, modelCourse, id, idhistory, c)

}
