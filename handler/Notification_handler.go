package handler

import (
	"balbibe/models"
	"balbibe/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNotification(c echo.Context) error {
	var notificationModel models.Notification

	err := c.Bind(&notificationModel)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return repository.CreateNotification(notificationModel, c)
}
