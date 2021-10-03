package database

import (
	"gits-echo-boilerplate/models"
	"testing"
)

var (
	idOrganisasi string = "qweqweqasdasd"
	idKelas      string = "asdasdqwexzcxz"
)

func TestCreateMahasiswa(t *testing.T) {
	Init()
	t.Run("create mahasiswa 1", func(t *testing.T) {
		data := models.Mahasiswa{
			Organisasi: idOrganisasi,
			Kelas:      idKelas,
			Nama:       "Jack Daniel",
		}
		m, err := CreateMahasiswa(&data)
		if err.Code > 0 {
			t.Fatal(err.Message)
		}
		if m.ID == "" {
			t.Fatal("Failed to create")
		}
	})
}
