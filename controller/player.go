package controller

import (
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

type PlayerController struct{}

func (pc *PlayerController) GetPlayersWithAccountID(accountID int) []models.Player {
	var player []models.Player
	if err := db.DB.Where("account_id = ?", accountID).Find(&player).Error; err != nil {
		return player
	}
	return player
}

func (pc *PlayerController) GetPropertiesPlayer(accountID, playerID int) bool {
	var count int64

	db.DB.Where("account_id = ? AND id = ?", accountID, playerID).Find(&models.Player{}).Count(&count)

	return count > 0
}
