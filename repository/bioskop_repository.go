package repository

import (
	"database/sql"

	"github.com/rendysp97/Api-Bioskop/database"
	"github.com/rendysp97/Api-Bioskop/model"
)

func CreateDataFromRepo(Db *sql.DB, newData *model.Bioskop) error {

	sqlStatement := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := Db.QueryRow(sqlStatement, newData.Nama, newData.Lokasi, newData.Rating).
		Scan(&newData.ID)

	return err
}

func GetDataBioskopRepo(Db *sql.DB, dataBioskop *[]model.Bioskop) error {

	rows, err := Db.Query("SELECT * FROM bioskop")

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data model.Bioskop

		err := rows.Scan(&data.ID, &data.Nama, &data.Lokasi, &data.Rating)

		if err != nil {
			return err
		}

		*dataBioskop = append(*dataBioskop, data)
	}

	return nil
}

func GetDetailBioskopRepo(DB *sql.DB, data *model.Bioskop, id int) error {

	err := database.Db.QueryRow("SELECT * FROM bioskop WHERE id = $1", id).Scan(&data.ID, &data.Nama, &data.Lokasi, &data.Rating)

	if err != nil {
		return err
	}
	return nil

}

func UpdateDataBioskopRepo(DB *sql.DB, data *model.Bioskop, id int) error {

	err := database.Db.QueryRow(" UPDATE bioskop SET nama = $1 , lokasi = $2 , rating = $3 WHERE id= $4 RETURNING id, nama, lokasi, rating", data.Nama, data.Lokasi, data.Rating, id).Scan(&data.ID, &data.Nama, &data.Lokasi, &data.Rating)

	if err != nil {
		return err
	}
	return nil
}

func DeleteDataBioskopRepo(DB *sql.DB, id int) error {

	database.Db.Exec("DELETE FROM bioskop WHERE id = $1", id)

	return nil
}
