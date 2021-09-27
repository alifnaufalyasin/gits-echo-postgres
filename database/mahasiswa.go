package database

import (
	"database/sql"
	"fmt"
	"gits-echo-boilerplate/models"
	"net/http"
)

type FilterSearch struct {
	Umur  int64
	Kelas string
}

func CreateMahasiswa(data *models.Mahasiswa) models.Error {
	query := `
		INSERT INTO
			mahasiswa(id, nama, umur, kelas, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6)
	`
	con := CreateCon()
	row := con.QueryRow(query,
		data.ID,
		data.Nama,
		data.Umur,
		data.Kelas,
		data.CreatedAt,
		data.UpdatedAt,
	)
	if row.Err() != nil {
		return models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to insert data",
		}
	}
	return models.Error{}
}

func GetMahasiswaByID(ID string) (m models.Mahasiswa, err models.Error) {
	query := `
	SELECT *
	FROM mahasiswa
	WHERE id = $1`

	con := CreateCon()
	row := con.QueryRow(query, ID)

	var (
		updatedAt sql.NullTime
		createdAt sql.NullTime
	)

	if err := row.Scan(
		&m.ID,
		&m.Nama,
		&m.Umur,
		&m.Kelas,
		&createdAt,
		&updatedAt,
	); err != nil {
		fmt.Println(err)
		return models.Mahasiswa{}, models.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	m.CreatedAt = createdAt.Time
	m.UpdatedAt = updatedAt.Time

	return
}

func GetAllMahasiswa(filter FilterSearch) ([]models.Mahasiswa, models.Error) {
	query := `
	SELECT *
	FROM mahasiswa
	`

	if filter.Kelas != "" && filter.Umur > 0 {
		query += fmt.Sprintf("WHERE kelas like '%%%s%%' AND umur = %d", filter.Kelas, filter.Umur)
	} else if filter.Kelas != "" {
		query += fmt.Sprintf("WHERE kelas like '%%%s%%'", filter.Kelas)
	} else if filter.Umur > 0 {
		query += fmt.Sprintf("WHERE umur = %d", filter.Umur)
	}

	con := CreateCon()
	var mahasiswa []models.Mahasiswa
	rows, err := con.Query(query)
	if err != nil {
		return []models.Mahasiswa{}, models.Error{
			Code:    500,
			Message: err.Error(),
		}
	}
	for rows.Next() {
		var m models.Mahasiswa
		if err := rows.Scan(&m.ID, &m.Nama, &m.Umur, &m.Kelas, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return mahasiswa, models.Error{
				Code:    500,
				Message: err.Error(),
			}
		}
		mahasiswa = append(mahasiswa, m)
	}
	if err = rows.Err(); err != nil {
		return mahasiswa, models.Error{
			Code:    500,
			Message: err.Error(),
		}
	}
	return mahasiswa, models.Error{}
}

func UpdateMahasiswa(data *models.Mahasiswa) (int64, models.Error) {
	query := `
	UPDATE mahasiswa
	SET
		nama = $2,
		umur = $3,
		kelas = $4,
		updated_at = $5
	WHERE id = $1
	`
	con := CreateCon()
	affectedRow, err := con.Exec(query,
		data.ID,
		data.Nama,
		data.Umur,
		data.Kelas,
		data.UpdatedAt,
	)
	if err != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	row, _ := affectedRow.RowsAffected()
	if row <= 0 {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengupdate data",
		}
	}

	return row, models.Error{}
}

func DeleteMahasiswa(id string) (int64, models.Error) {
	query := `
	DELETE FROM mahasiswa
	WHERE id = $1
	`
	con := CreateCon()
	affectedRow, err := con.Exec(query, id)

	if err != nil {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	row, _ := affectedRow.RowsAffected()
	if row <= 0 {
		return 0, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menghapus data",
		}
	}

	return row, models.Error{}
}
