package config

import (
	"eulabs/src/main/core/product"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "localhost"
	user := "mysql"
	password := "123456"
	dbName := "eulabs"
	port := 3306

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		log.Panic(e)
	}

	err := database.AutoMigrate(&product.Product{})

	if err != nil {
		log.Panic(e)
	}

	fmt.Println("Database connected!")

}

func DB() *gorm.DB {
	return database
}
