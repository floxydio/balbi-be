package handler

import (
	"balbibe/models"
	"balbibe/repository"

	"github.com/labstack/echo/v4"
)


func GetAllBlog(c echo.Context) error {
	var infoModel []models.InfoBlog
	return repository.GetListBlog(infoModel,c)
}