package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error

	dsn := "root:own3306@tcp(server.pca7.com:3306)/cigin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return DB
}
