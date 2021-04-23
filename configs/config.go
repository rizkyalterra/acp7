package configs

import (
	"fmt"
	"project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USER = "root"
const DB_PASS = "123ABC4d."
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"
const DB_NAME = "acp7"

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
	// DB.AutoMigrate(&models.School{})
	// DB.AutoMigrate(&models.UserCompany{})
	// DB.AutoMigrate(&models.Company{})
}
