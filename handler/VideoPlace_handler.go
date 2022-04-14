package handler

import (
	"balbibe/models"
	"balbibe/repository"

	"github.com/labstack/echo/v4"
)

func GetVideoPlaceAll(c echo.Context) error {
	var videoPlace []models.VideoPlace
	return repository.GetVideoPlaceAll(c, videoPlace)
}
