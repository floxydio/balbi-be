package handler

import (
	"balbibe/database"
	models "balbibe/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllKomunitas(c echo.Context) error {
	komunitasModel := []models.Komunitas{}

	err := database.DB.Find(&komunitasModel).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Message: "Successfully retrieved all komunitas",
		Data:    komunitasModel,
	})
}

func GetAllPostById(c echo.Context) error {
	type GetPostbyId struct {
		Id            uint   `json:"id"`
		Nama          string `json:"nama"`
		Title         string `json:"title"`
		NamaKomunitas string `json:"nama_komunitas"`
	}

	var product []GetPostbyId

	id := c.Param("id")

	err := database.DB.Raw("SELECT komunitas_post.id,user.nama, komunitas_post.title,komunitas_post.content,komunitas.nama FROM komunitas_post LEFT JOIN user ON komunitas_post.created_by = user.id LEFT JOIN komunitas ON komunitas_post.id = komunitas.id WHERE komunitas_post.id = ?", id).Scan(&product).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  http.StatusOK,
		Message: "Successfully retrieved all post by id",
		Data:    product,
	})

}
