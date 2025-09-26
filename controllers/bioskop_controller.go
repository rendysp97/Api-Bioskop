package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rendysp97/Api-Bioskop/database"
	"github.com/rendysp97/Api-Bioskop/model"
	"github.com/rendysp97/Api-Bioskop/repository"
)

func CreateDataBioskop(ctx *gin.Context) {
	var newData model.Bioskop

	if err := ctx.ShouldBindJSON(&newData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newData.Nama == "" || newData.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	err := repository.CreateDataFromRepo(database.Db, &newData)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    newData,
	})

}

func GetAllData(ctx *gin.Context) {

	var dataBioskop []model.Bioskop

	repository.GetDataBioskopRepo(database.Db, &dataBioskop)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    dataBioskop,
		"message": "Success",
	})

}

func GetDetailBioskop(ctx *gin.Context) {

	fromParams := ctx.Param("id")

	id, err := strconv.Atoi(fromParams)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Need Valid Id "})
		return
	}

	var data model.Bioskop

	repository.GetDetailBioskopRepo(database.Db, &data, id)

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func UpdateDataBioskop(ctx *gin.Context) {

	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	var data model.Bioskop

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.Nama == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama tidak boleh kosong"})
		return
	}

	if data.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Lokasi tidak boleh kosong"})
		return
	}

	if data.Rating < 0 || data.Rating > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Rating  tidak boleh 0 dan lebih dari 100"})
		return
	}

	repository.UpdateDataBioskopRepo(database.Db, &data, id)

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Success Update",
	})

}

func DeleteDataBioskop(ctx *gin.Context) {

	paramsId := ctx.Param("id")

	id, err := strconv.Atoi(paramsId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	repository.DeleteDataBioskopRepo(database.Db, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete",
	})

}
