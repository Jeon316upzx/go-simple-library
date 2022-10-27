package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Host     = "127.0.0.1"
	User     = "root"
	Password = "*******"
	DBName   = "library"
	Port     = "3306"
)

var DB *gorm.DB

func Setup() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		User,
		Host,
		Port,
		DBName,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
