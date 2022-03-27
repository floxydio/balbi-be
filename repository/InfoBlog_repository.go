package repository

import (
	"balbibe/database"
	"balbibe/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetListBlog(listinfo []models.InfoBlog, c echo.Context) error {
	err := database.DB.Find(listinfo).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data":    listinfo,
		"message": "Successfully Get List Info Blog",
	})
}
