package models

import (
	"time"
)

const (
	MataKuliahTableName = "mata_kuliah"
)

type MataKuliah struct {
	ID        string    `gorm:"type:varchar(50);primary_key" json:"id"`
	Nama      string    `gorm:"type:varchar(100);not_null" json:"nama"`
	Kode      string    `gorm:"type:varchar(10);not_null" json:"kode"`
	Nilai     []Nilai   `gorm:"foreignKey:MataKuliah" json:"nilai"`
	CreatedAt time.Time `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not_null" json:"updated_at"`
}

// TableName specifies table name for MataKuliahModel.
func (model *MataKuliah) TableName() string {
	return MataKuliahTableName
}
