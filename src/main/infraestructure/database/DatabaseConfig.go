package config

import (
	"eulabs/src/main/core/product"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var database *gorm.DB
var e error

func DatabaseInit() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
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
