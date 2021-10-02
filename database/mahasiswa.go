package database

import (
	"gits-echo-boilerplate/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type FilterSearch struct {
	Umur  int64
	Kelas string
}

var (
	mahasiswa models.Mahasiswa
)

func CreateMahasiswa(data *models.Mahasiswa) (models.Mahasiswa, models.Error) {
	db := CreateCon()
	res := db.Create(data)
	if res.Error != nil {
		return models.Mahasiswa{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	if res.RowsAffected <= 0 {
		return models.Mahasiswa{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menambahkan data",
		}
	}
	return *data, models.Error{}
}

func GetMahasiswaByID(ID string) (models.Mahasiswa, models.Error) {
	db := CreateCon()
	// result := map[string]interface{}{}
	res := db.First(&mahasiswa, "id = ?", ID)
	if res.Error != nil {
		return models.Mahasiswa{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	var m models.Mahasiswa
	err := res.Scan(&m)
	if err.Error != nil {
		return models.Mahasiswa{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	return m, models.Error{}
}

func GetAllMahasiswa(filter FilterSearch) ([]models.Mahasiswa, models.Error) {
	var mahasiswas []models.Mahasiswa
	db := CreateCon()
	var res *gorm.DB
	if filter.Kelas != "" && filter.Umur > 0 {
		res = db.Where("kelas LIKE ? AND umur = ?", filter.Kelas, filter.Umur).Find(&mahasiswas)
	} else if filter.Kelas != "" {
		res = db.Where("kelas LIKE ?", filter.Kelas).Find(&mahasiswas)
	} else if filter.Umur > 0 {
		res = db.Where("umur = ?", filter.Umur).Find(&mahasiswas)
	} else {
		res = db.Find(&mahasiswas)
	}

	if res.Error != nil {
		return []models.Mahasiswa{}, models.Error{
			Code:    500,
			Message: res.Error.Error(),
		}
	}
	return mahasiswas, models.Error{}
}

func UpdateMahasiswa(data *models.Mahasiswa) (int64, models.Error) {
	m := models.Mahasiswa{}
	db := CreateCon()
	err := db.First(&m, "id = ?", data.ID)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	m.Nama = data.Nama
	m.Kelas = data.Kelas
	m.Organisasi = data.Organisasi
	m.UpdatedAt = time.Now()

	err = db.Save(&m)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}

func DeleteMahasiswa(id string) (int64, models.Error) {
	db := CreateCon()
	err := db.Delete(&models.Mahasiswa{}, "id = ?", id)
	return err.RowsAffected, models.Error{}
}
