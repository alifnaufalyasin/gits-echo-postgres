package models

import "time"

const (
	OrganisasiTableName = "organisasi"
)

type Organisasi struct {
	ID        string      `gorm:"type:varchar(50);primary_key" json:"id"`
	Mahasiswa []Mahasiswa `gorm:"foreignKey:Organisasi" json:"mahasiswa"`
	Nama      string      `gorm:"type:varchar(100);not_null" json:"nama"`
	Deskripsi string      `gorm:"type:varchar(500);not_null" json:"deskripsi"`
	Lokasi    string      `gorm:"type:varchar(100);not_null" json:"lokasi"`
	CreatedAt time.Time   `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt time.Time   `gorm:"type:timestamptz;not_null" json:"updated_at"`
}

// TableName specifies table name for KelasModel.
func (model *Organisasi) TableName() string {
	return OrganisasiTableName
}
