package models

import "time"

const (
	NilaiTableName = "nilai"
)

type Nilai struct {
	ID         string    `gorm:"type:varchar(50);primary_key" json:"id"`
	Angka      string    `gorm:"type:varchar(5);not_null" json:"angka"`
	Mahasiswa  string    `gorm:"type:varchar(50);not_null" json:"mahasiswa"`
	MataKuliah string    `gorm:"type:varchar(50);not_null" json:"mata_kuliah"`
	CreatedAt  time.Time `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamptz;not_null" json:"updated_at"`
}

// TableName specifies table name for NilaiModel.
func (model *Nilai) TableName() string {
	return NilaiTableName
}
