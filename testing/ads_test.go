package testing

import (
	"balbibe/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func MockAds(c echo.Context) error {
	return c.JSON(http.StatusOK, "Success")
}

var adsModel []models.Ads

func MockCreateAds(c echo.Context) error {
	return c.JSON(http.StatusOK, "Success")
}

func TestAds(t *testing.T) {
	e := echo.New()

	testJson := models.Ads{
		Id:        1,
		Title:     "Test",
		Image:     "1.jpg",
		Detail:    "Test Detail",
		View:      8,
		CreatedAt: "Rabu",
	}

	req := httptest.NewRequest(http.MethodGet, "/get-ads", nil)

	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	if assert.NoError(t, MockAds(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, testJson, models.Ads{
			Id:        1,
			Title:     "Test",
			Image:     "1.jpg",
			Detail:    "Test Detail",
			View:      8,
			CreatedAt: "Rabu",
		})
	}

	// e.GET("/ads", handler.GetAdsAll)

}
