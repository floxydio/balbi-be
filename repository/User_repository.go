package repository

import (
	"balbibe/database"
	"balbibe/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User, c echo.Context) error {
	err := database.DB.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		// "data": userModel,
		"message": "Successfully Created",
	})
}

func LoginUser(username string, user models.User, c echo.Context) error {
	errs := database.DB.Where("email = ?", &username).First(&user).Error

	if errs != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errs.Error(),
		})
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.FormValue("password")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Password not match",
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 10080).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid JWT",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Sign In",
		"token":   t,
		"data":    user,
	})
}

func EditUserById(id int64, user models.User, c echo.Context) error {
	err := database.DB.Model(&user).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Updated",
	})
}

func VerifyEmail(user string, token string, c echo.Context) error {
	err := database.DB.Model(models.User{}).Where("email = ?", user).Update("email_valid", "1").Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Email Verification berhasil",
	})
}
