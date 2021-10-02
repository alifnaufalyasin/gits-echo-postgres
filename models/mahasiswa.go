package models

import (
	"time"
)

const (
	MahasiswaTableName = "mahasiswa"
)

type Mahasiswa struct {
	ID         string    `gorm:"type:varchar(50);primary_key" json:"id"`
	Nama       string    `gorm:"type:varchar(100);not_null" json:"nama"`
	Kelas      string    `gorm:"type:varchar(50);not_null" json:"kelas"`
	Nilai      []Nilai   `gorm:"foreignKey:Mahasiswa" json:"nilai"`
	Organisasi string    `gorm:"type:varchar(50);not_null" json:"organisasi"`
	CreatedAt  time.Time `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamptz;not_null" json:"updated_at"`
}

// TableName specifies table name for MahasiswaModel.
func (model *Mahasiswa) TableName() string {
	return MahasiswaTableName
}
