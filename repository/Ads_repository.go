package repository

import (
	"balbibe/database"
	"balbibe/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetRepoAds(ads []models.Ads, c echo.Context) error {
	err := database.DB.Find(&ads).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Get Ads",
		"data":    ads,
	})
}

func CountInView(ads models.Ads, param string, c echo.Context) error {
	err := database.DB.Model(&ads).Where("id = ?", param).Update("view", gorm.Expr("view + ?", 1)).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Count In View",
	})
}
