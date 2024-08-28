package db

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Players{})
	DB.AutoMigrate(&models.NewsTicket{})
	DB.AutoMigrate(&models.Towns{})
	DB.AutoMigrate(&models.PlayersOnline{})
	utils.Info("Update MySQL")
}

func ConnectionMysql(user, password, host, nameDB string, port uint16, debugMode bool) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, nameDB)

	logDebug := logger.Silent
	if debugMode {
		logDebug = logger.Warn
	}
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logDebug),
	})
	if err != nil {
		return err
	}
	utils.Info("Connection MySQL")
	AutoMigrate(DB)

	return nil
}
