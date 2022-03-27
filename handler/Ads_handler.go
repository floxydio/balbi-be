package handler

import (
	"balbibe/models"
	"balbibe/repository"

	"github.com/labstack/echo/v4"
)

func GetAdsAll(c echo.Context) error {
	var adsModel []models.Ads

	return repository.GetRepoAds(adsModel, c)

}

func CountInView(c echo.Context) error {
	var adsModel models.Ads
	param := c.Param("id")

	return repository.CountInView(adsModel, param, c)
}