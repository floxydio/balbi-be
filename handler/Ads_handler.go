package handler

import (
	"balbibe/models"
	"balbibe/repository"
	"io"
	"net/http"
	"os"

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

func CreateAds(c echo.Context) error {
	adsModel := models.Ads{}
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create("uploads/ads/" + file.Filename)

	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	errs := c.Bind(&adsModel)
	if errs != nil {
		return errs
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Create Ads",
	})

}
