package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simpleLibrary/config"
	"github.com/simpleLibrary/models"
)

func Get_all_books(ctx *gin.Context) {
	books := []models.Book{}
	config.DB.Find(&books)
	ctx.JSON(http.StatusOK, &books)
}

func Create_book(ctx *gin.Context) {
	var book models.Book
	ctx.BindJSON(&book)
	config.DB.Create(&book)
	ctx.JSON(http.StatusOK, &book)
}

func Get_book(ctx *gin.Context) {
	var book models.Book
	config.DB.Where("id=?", ctx.Param("id")).First(&book)
	ctx.JSON(http.StatusOK, &book)

}

func Get_book_by_name(ctx *gin.Context) {
	var book models.Book
	config.DB.Where("title = ?", ctx.Param("title")).First(&book)
	ctx.JSON(http.StatusOK, &book)
}

func Delete_book(ctx *gin.Context) {
	var book models.Book
	config.DB.Where("id =?", ctx.Param("id")).Delete(&book)
	ctx.JSON(http.StatusOK, &book)
}
