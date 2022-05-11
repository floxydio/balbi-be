package handler

import (
	"balbibe/database"
	"balbibe/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateArticle(c echo.Context) error {
	article := new(models.Article)
	if err := c.Bind(article); err != nil {
		return err
	}
	if err := database.DB.Create(&article).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   &article,
	})
}

func GetArticle(c echo.Context) error {
	var article []models.Article

	if err := database.DB.Find(&article).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   article,
	})

}
