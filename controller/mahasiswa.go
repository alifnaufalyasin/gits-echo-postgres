package controller

import (
	"gits-echo-boilerplate/database"
	"gits-echo-boilerplate/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid"
)

func CreateMahasiswa(c echo.Context) error {
	m := new(models.Mahasiswa)
	if err := c.Bind(m); err != nil {
		return (models.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}).Response(c)
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	m.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	maha, err := database.CreateMahasiswa(m)
	if err.Code > 0 {
		return (models.JSONResponse{
			Code:    err.Code,
			Message: "Gagal menambahkan data mahasiswa",
		}).Response(c)
	}
	resp := models.JSONResponseData{
		Code:    http.StatusCreated,
		Data:    maha,
		Message: "Data Created",
	}
	return resp.Response(c)

}

func GetMahasiswaByID(c echo.Context) error {
	id := c.Param("id")

	m, err := database.GetMahasiswaByID(id)
	if err.Code > 0 {
		return err.Response(c)
	}

	return (models.JSONResponseData{
		Code:    http.StatusOK,
		Data:    m,
		Message: "Success",
	}).Response(c)
}

func GetAllMahasiswa(c echo.Context) error {
	kelas := c.QueryParam("kelas")
	umur := c.QueryParam("umur")
	i, _ := strconv.ParseInt(umur, 10, 64)
	m, err := database.GetAllMahasiswa(database.FilterSearch{
		Umur:  i,
		Kelas: kelas,
	})
	if err.Code > 0 {
		return err.Response(c)
	}

	return (models.JSONResponseData{
		Code:    http.StatusOK,
		Data:    m,
		Message: "Success",
	}).Response(c)

}

func UpdateMahasiswa(c echo.Context) error {
	id := c.Param("id")

	m := new(models.Mahasiswa)
	if err := c.Bind(m); err != nil {
		return (models.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}).Response(c)
	}
	m.ID = id
	m.UpdatedAt = time.Now()

	_, err := database.UpdateMahasiswa(m)
	if err.Code > 0 {
		return err.Response(c)
	}
	resp := models.JSONResponseData{
		Code:    http.StatusOK,
		Data:    m,
		Message: "Success",
	}
	return resp.Response(c)
}

func DeleteMahasiswa(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteMahasiswa(id)

	if err.Code > 0 {
		return err.Response(c)
	}

	resp := models.JSONResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}
	return resp.Response(c)
}
