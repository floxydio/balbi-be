package main

import (
	"balbibe/database"
	"balbibe/models"
	"balbibe/routes"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func msgForValidate(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field is required"
	case "email":
		return "Invalid email"
	case "min=6":
		return "Minimum 6 characters"
	}
	return fe.Error() // default error
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = models.ApiError{fe.Field(), msgForValidate(fe)}
			}
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"errors": out})
		}
	}
	return nil
}

func main() {
	// Echo instance
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	database.Connect()
	// Middleware
	e.Use(middleware.Recover())
	// Routes
	routes.Setup(e)

	// Start server
	e.Logger.Fatal(e.Start(":2000"))
}
