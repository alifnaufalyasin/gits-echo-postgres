package database

import (
	"gits-echo-boilerplate/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func CreateKelas(data *models.Kelas) (models.Kelas, models.Error) {
	db := CreateCon()
	res := db.Create(data)
	if res.Error != nil {
		return models.Kelas{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: res.Error.Error(),
		}
	}
	if res.RowsAffected <= 0 {
		return models.Kelas{}, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menambahkan data",
		}
	}
	return *data, models.Error{}
}

// func GetMahasiswaByID(ID string) (models.Mahasiswa, models.Error) {
// 	db := CreateCon()
// 	// result := map[string]interface{}{}
// 	res := db.First(&mahasiswa, "id = ?", ID)
// 	if res.Error != nil {
// 		return models.Mahasiswa{}, models.Error{
// 			Code:    http.StatusInternalServerError,
// 			Message: res.Error.Error(),
// 		}
// 	}
// 	var m models.Mahasiswa
// 	err := res.Scan(&m)
// 	if err.Error != nil {
// 		return models.Mahasiswa{}, models.Error{
// 			Code:    http.StatusInternalServerError,
// 			Message: res.Error.Error(),
// 		}
// 	}
// 	return m, models.Error{}
// }

func GetAllKelas(kelas string) ([]models.Kelas, models.Error) {
	var AllKelas []models.Kelas
	db := CreateCon()
	var res *gorm.DB
	if kelas != "" {
		res = db.Where("kelas LIKE ?", kelas).Find(&AllKelas)
	} else {
		res = db.Find(&AllKelas)
	}

	if res.Error != nil {
		return []models.Kelas{}, models.Error{
			Code:    500,
			Message: res.Error.Error(),
		}
	}
	return AllKelas, models.Error{}
}

func UpdateKelas(data *models.Kelas) (int64, models.Error) {
	k := models.Kelas{}
	db := CreateCon()
	err := db.First(&k, "id = ?", data.ID)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	k.Nama = data.Nama
	k.Mahasiswa = data.Mahasiswa
	k.WaliDosen = data.WaliDosen
	k.UpdatedAt = time.Now()

	err = db.Save(&k)
	if err.Error != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error.Error(),
		}
	}
	return err.RowsAffected, models.Error{}
}

func DeleteKelas(id string) (int64, models.Error) {
	db := CreateCon()
	err := db.Delete(&models.Kelas{}, id)
	return err.RowsAffected, models.Error{}
}
