package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/simpleLibrary/config"
	"github.com/simpleLibrary/models"
	"github.com/simpleLibrary/routes"
)

func main() {
	router := gin.New()
	db, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	fmt.Println("Successfully connected to database")

	err1 := db.AutoMigrate(models.Book{})

	if err1 != nil {
		log.Fatal("Migration error ", err1)
	}

	fmt.Println("Successfully migrated model to database")

	routes.BooksRoute(router)
	router.Run(":8080")

}
