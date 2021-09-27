package routes

import (
	"gits-echo-boilerplate/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/mahasiswa", controller.CreateMahasiswa)
	e.GET("/mahasiswa/:id", controller.GetMahasiswaByID)
	e.GET("/mahasiswa", controller.GetAllMahasiswa)
	e.PUT("/mahasiswa/:id", controller.UpdateMahasiswa)
	e.DELETE("/mahasiswa/:id", controller.DeleteMahasiswa)

	// e.Logger.Fatal(e.Start(":4132"))
	return e
}
