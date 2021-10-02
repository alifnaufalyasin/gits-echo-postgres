package migration

import (
	"fmt"
	"gits-echo-boilerplate/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.Migrator().CreateTable(&models.Kelas{}, &models.Organisasi{}, &models.Mahasiswa{}, &models.MataKuliah{}, &models.Nilai{})
	if err != nil {
		fmt.Println(err)
	}
}
