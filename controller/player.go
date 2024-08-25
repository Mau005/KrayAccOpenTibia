package controller

import (
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

type PlayerController struct{}

func (pc *PlayerController) GetPlayersWithAccountID(accountID int) []models.Players {
	var player []models.Players
	if err := db.DB.Where("account_id = ?", accountID).Find(&player).Error; err != nil {
		return player
	}
	return player
}

func (pc *PlayerController) GetPropertiesPlayer(accountID, playerID int) bool {
	var count int64

	db.DB.Where("account_id = ? AND id = ?", accountID, playerID).Find(&models.Players{}).Count(&count)

	return count > 0
}

func (pc *PlayerController) GetPlayerLimits(count int) (player []models.Players) {
	db.DB.Find(&player).Limit(count).Order("asc level")
	return
}
