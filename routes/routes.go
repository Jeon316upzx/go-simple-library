package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/simpleLibrary/controllers"
)

func BooksRoute(router *gin.Engine) {
	router.GET("/books", controller.Get_all_books)
	router.POST("/book", controller.Create_book)
	router.DELETE("/book/:id", controller.Delete_book)
	router.GET("/book/:id", controller.Get_book)
	router.GET("/book/name/:title", controller.Get_book_by_name)
}
