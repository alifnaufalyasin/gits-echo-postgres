package database

import (
	"gits-echo-boilerplate/models"
	"testing"
)

func TestCreateKelas(t *testing.T) {
	t.Run("create kelas 1", func(t *testing.T) {
		dataKelas := models.Kelas{
			Nama:      "IF-42-09",
			WaliDosen: "Pak Nam Do San",
		}
		k, err := CreateKelas(&dataKelas)
		if err.Code > 0 {
			t.Fatal(err)
		}
		if k.ID == "" {
			t.Fatal("gagal create data")
		}
	})
}
