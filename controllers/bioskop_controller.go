package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rendysp97/Api-Bioskop/database"
)

type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

func CreateDataBioskop(ctx *gin.Context) {
	var newData Bioskop

	if err := ctx.ShouldBindJSON(&newData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newData.Nama == "" || newData.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	sqlStatement := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := database.Db.QueryRow(sqlStatement, newData.Nama, newData.Lokasi, newData.Rating).
		Scan(&newData.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    newData,
	})

}
