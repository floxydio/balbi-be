package repository

import (
	"balbibe/database"
	"balbibe/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePostAnakPertama(c echo.Context, anakPertama models.DaftarAnak) error {
	err := database.DB.Create(&anakPertama).Error

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   &anakPertama,
	})
}

func CreateAnakKedua(c echo.Context, anakKedua models.DaftarAnak) error {
	err := database.DB.Create(&anakKedua).Error

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   &anakKedua,
	})
}

func CreateAnakKetiga(c echo.Context, anakKetiga models.DaftarAnak) error {
	err := database.DB.Create(&anakKetiga).Error

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data":   &anakKetiga,
	})
}

func GetAnakById(c echo.Context, id string) error {
	var anak []models.DaftarAnak
	var anakKedua []models.DaftarAnakKedua
	var anakKetiga []models.DaftarAnakKetiga
	sqlQuery := fmt.Sprintf("SELECT usr.nama,usr.email,a.gender,a.namaanak,a.nama_panggilan,a.tanggal_lahir,a.panjang_badan,a.lingkar_kepala FROM anak a LEFT JOIN user usr ON a.userid = usr.id WHERE usr.id = %s", id)
	sqlQueryKedua := fmt.Sprintf("SELECT usr.nama,usr.email,a.genderkedua,a.nama_anak_kedua,a.nama_panggilan_kedua,a.tanggal_lahir_kedua,a.panjang_badan_kedua,a.lingkar_kepala_kedua, a.berat_badan_kedua FROM anakkedua a LEFT JOIN user usr ON a.userid = usr.id  WHERE usr.id = %s", id)
	sqlQueryKetiga := fmt.Sprintf("SELECT usr.nama,usr.email,a.genderketiga,a.nama_anak_ketiga,a.nama_panggilan_ketiga,a.tanggal_lahir_ketiga,a.panjang_badan_ketiga,a.lingkar_kepala_ketiga, a.berat_badan_ketiga FROM anakketiga a LEFT JOIN user usr ON a.userid = usr.id  WHERE usr.id = %s", id)
	err := database.DB.Raw(sqlQuery).Scan(&anak).Error
	errTwo := database.DB.Raw(sqlQueryKedua).Scan(&anakKedua).Error
	errThree := database.DB.Raw(sqlQueryKetiga).Scan(&anakKetiga).Error

	if err != nil {
		return err
	}

	if errTwo != nil {
		return errTwo
	}
	if errThree != nil {
		return errThree
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"anak":       anak,
			"anakKedua":  anakKedua,
			"anakKetiga": anakKetiga,
		},
	})
}
