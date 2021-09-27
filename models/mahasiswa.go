package models

import (
	"net/http"
	"time"
)

type Mahasiswa struct {
	ID        string    `json:"id"`
	Nama      string    `json:"nama"`
	Umur      int64     `json:"umur"`
	Kelas     string    `json:"kelas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Mahasiswa) Validate() Error {
	if m.ID == "" {
		return Error{
			Code:    http.StatusBadRequest,
			Message: "Mahasiswa ID not found",
		}
	}
	if m.Nama == "" {
		return Error{
			Code:    http.StatusBadRequest,
			Message: "Mahasiswa Name not found",
		}
	}
	if m.Umur <= 0 {
		return Error{
			Code:    http.StatusBadRequest,
			Message: "Mahasiswa Age not found",
		}
	}
	if m.Kelas == "" {
		return Error{
			Code:    http.StatusBadRequest,
			Message: "Mahasiswa Class not found",
		}
	}
	return Error{}
}
