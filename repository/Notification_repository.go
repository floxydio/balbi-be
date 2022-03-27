package repository

import (
	"balbibe/database"
	"balbibe/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNotification(notificationModel models.Notification, c echo.Context) error {
	err := database.DB.Create(&notificationModel).Error

	if err != nil {
		return c.JSON(http.StatusNotImplemented, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"data":    &notificationModel,
		"message": "Successfully Created Notification",
	})
}
