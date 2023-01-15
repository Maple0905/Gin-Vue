package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DBDriver := os.Getenv("DB_DRIVER")
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

	DB, err = gorm.Open(DBDriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database", DBDriver)
		log.Fatal("connection error: ", err)
	} else {
		fmt.Println("We are connected to the database", DBDriver)
	}

	DB.AutoMigrate(&User{}, &Role{})
}
