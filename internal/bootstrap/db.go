package bootstrap

import (
	"time"

	"cigin/pkg/model"

	"gorm.io/gorm"
)

// SetupDB DB初始化
func SetupDB() *gorm.DB {
	db := model.ConnectDB()
	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(100)

	if err != nil {
		panic("数据库链接失败")
	}

	return db
}
