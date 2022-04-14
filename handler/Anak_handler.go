package handler

import (
	"balbibe/models"
	"balbibe/repository"

	"github.com/labstack/echo/v4"
)

func CreateAnakPertama(c echo.Context) error {
	var anakPertama models.DaftarAnak

	if err := c.Bind(&anakPertama); err != nil {
		return err
	}
	return repository.CreatePostAnakPertama(c, anakPertama)
}

func CreateAnakKedua(c echo.Context) error {
	var anakKedua models.DaftarAnak

	if err := c.Bind(&anakKedua); err != nil {
		return err
	}
	return repository.CreateAnakKedua(c, anakKedua)
}

func CreateAnakKetiga(c echo.Context) error {
	var anakKetiga models.DaftarAnak

	if err := c.Bind(&anakKetiga); err != nil {
		return err
	}
	return repository.CreateAnakKetiga(c, anakKetiga)
}

func GetAnakById(c echo.Context) error {
	id := c.Param("id")
	return repository.GetAnakById(c, id)
}
