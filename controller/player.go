package controller

import (
	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
)

type PlayerController struct{}

func (pc *PlayerController) GetPlayerWithAccountID(accountID uint) []models.Player {
	var player []models.Player
	if err := db.DB.Where("account_id = ?", accountID).Find(&player).Error; err != nil {
		return player
	}
	return player
}
