package models

import "time"

const (
	KelasTableName = "kelas"
)

type Kelas struct {
	ID        string      `gorm:"type:varchar(50);primary_key" json:"id"`
	Nama      string      `gorm:"type:varchar(50);not_null" json:"nama"`
	Mahasiswa []Mahasiswa `gorm:"foreignKey:Kelas" json:"mahasiswa"`
	WaliDosen string      `gorm:"type:varchar(50);not_null" json:"wali_dosen"`
	CreatedAt time.Time   `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt time.Time   `gorm:"type:timestamptz;not_null" json:"updated_at"`
}

// TableName specifies table name for KelasModel.
func (model *Kelas) TableName() string {
	return KelasTableName
}
