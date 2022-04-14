package repository

import (
	"balbibe/database"
	"balbibe/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetVideoPlaceAll(c echo.Context, videoPlace []models.VideoPlace) error {
	err := database.DB.Find(&videoPlace).Error

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   &videoPlace,
	})
}

