package database

import (
	"gits-echo-boilerplate/models"
	"net/http"
	"time"
)

var (
	nilai models.Nilai
)

func CreateNilai(data *models.Nilai) (models.Nilai, models.Error) {
	db := CreateCon()
	res := db.Create(data)
	if res.Error != nil {
		return models.Nilai{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	if res.RowsAffected <= 0 {
		return models.Nilai{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menambahkan data",
		}
	}
	return *data, models.Error{}
}

func GetNilaiByID(ID string) (models.Nilai, models.Error) {
	db := CreateCon()
	// result := map[string]interface{}{}
	res := db.First(&nilai, "id = ?", ID)
	if res.Error != nil {
		return models.Nilai{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	var m models.Nilai
	err := res.Scan(&m)
	if err.Error != nil {
		return models.Nilai{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	return m, models.Error{}
}

func GetAllNilai(organisasi string) ([]models.Nilai, models.Error) {
	var AllNilai []models.Nilai
	db := CreateCon()
	// var res *gorm.DB

	res := db.Find(&AllNilai)

	if res.Error != nil {
		return []models.Nilai{}, models.Error{
			Code:    500,
			Message: res.Error.Error(),
		}
	}
	return AllNilai, models.Error{}
}

func UpdateNilai(data *models.Nilai) (int64, models.Error) {
	n := models.Nilai{}
	db := CreateCon()
	err := db.First(n, "id = ?", data.ID)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}

	n.Angka = data.Angka
	n.UpdatedAt = time.Now()

	err = db.Save(&n)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}

func DeleteNilai(id string) (int64, models.Error) {
	db := CreateCon()
	err := db.Delete(&models.Nilai{}, id)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}
