package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Host     = "mysqldb"
	User     = "root"
	Password = "library1818"
	DBName   = "library"
	Port     = "3306"
)

var DB *gorm.DB

func Setup() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		User,
		Password,
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
