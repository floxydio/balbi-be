package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "MAU NGAPAAAINN :P",
	})
}
