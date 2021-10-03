package database

import (
	"gits-echo-boilerplate/models"
	"net/http"
	"time"
)

var (
	mataKuliah models.MataKuliah
)

func CreateMataKuliah(data *models.MataKuliah) (models.MataKuliah, models.Error) {
	db := CreateCon()
	res := db.Create(data)
	if res.Error != nil {
		return models.MataKuliah{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	if res.RowsAffected <= 0 {
		return models.MataKuliah{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menambahkan data",
		}
	}
	return *data, models.Error{}
}

func GetMataKuliahByID(ID string) (models.MataKuliah, models.Error) {
	db := CreateCon()
	// result := map[string]interface{}{}
	res := db.First(&mataKuliah, "id = ?", ID)
	if res.Error != nil {
		return models.MataKuliah{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	var m models.MataKuliah
	err := res.Scan(&m)
	if err.Error != nil {
		return models.MataKuliah{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	return m, models.Error{}
}

func GetAllMataKuliah(organisasi string) ([]models.MataKuliah, models.Error) {
	var AllMataKuliah []models.MataKuliah
	db := CreateCon()
	// var res *gorm.DB

	res := db.Find(&AllMataKuliah)

	if res.Error != nil {
		return []models.MataKuliah{}, models.Error{
			Code:    500,
			Message: res.Error.Error(),
		}
	}
	return AllMataKuliah, models.Error{}
}

func UpdateMataKuliah(data *models.MataKuliah) (int64, models.Error) {
	mk := models.MataKuliah{}
	db := CreateCon()
	err := db.First(mk, "id = ?", data.ID)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}

	mk.Kode = data.Kode
	mk.Nama = data.Nama
	mk.UpdatedAt = time.Now()

	err = db.Save(&mk)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}

func DeleteMataKuliah(id string) (int64, models.Error) {
	db := CreateCon()
	err := db.Delete(&models.MataKuliah{}, id)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}
