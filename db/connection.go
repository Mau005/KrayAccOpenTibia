package db

import (
	"fmt"
	"log"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Player{})
}

func ConnectionMysql(user, password, host, name string, port int, debugMode bool) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name)

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return nil, err
	}
	AutoMigrate(DB)
	log.Println("connection MYSQL ok")
	return DB, nil
}
