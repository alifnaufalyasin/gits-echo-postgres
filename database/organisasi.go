package database

import (
	"gits-echo-boilerplate/models"
	"net/http"
	"time"
)

var (
	organisasi models.Organisasi
)

func CreateOrganisasi(data *models.Organisasi) (models.Organisasi, models.Error) {
	db := CreateCon()
	res := db.Create(data)
	if res.Error != nil {
		return models.Organisasi{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	if res.RowsAffected <= 0 {
		return models.Organisasi{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menambahkan data",
		}
	}
	return *data, models.Error{}
}

func GetOrganisasiByID(ID string) (models.Organisasi, models.Error) {
	db := CreateCon()
	// result := map[string]interface{}{}
	res := db.First(&organisasi, "id = ?", ID)
	if res.Error != nil {
		return models.Organisasi{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	var m models.Organisasi
	err := res.Scan(&m)
	if err.Error != nil {
		return models.Organisasi{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	return m, models.Error{}
}

func GetAllOrganisasi(organisasi string) ([]models.Organisasi, models.Error) {
	var AllOrganisasi []models.Organisasi
	db := CreateCon()
	// var res *gorm.DB

	res := db.Find(&AllOrganisasi)

	if res.Error != nil {
		return []models.Organisasi{}, models.Error{
			Code:    500,
			Message: res.Error.Error(),
		}
	}
	return AllOrganisasi, models.Error{}
}

func UpdateOrganisasi(data *models.Organisasi) (int64, models.Error) {
	o := models.Organisasi{}
	db := CreateCon()
	err := db.First(&o, "id = ?", data.ID)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	o.Nama = data.Nama
	o.Mahasiswa = data.Mahasiswa
	o.Deskripsi = data.Deskripsi
	o.Lokasi = data.Lokasi
	o.UpdatedAt = time.Now()

	err = db.Save(&o)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}

func DeleteOrganisasi(id string) (int64, models.Error) {
	db := CreateCon()
	err := db.Delete(&models.Organisasi{}, id)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}
