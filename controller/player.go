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
	db.DB.Order("level desc").Limit(count).Find(&player)
	return
}

func (pc *PlayerController) CreatePlayer(player models.Players) (models.Players, error) {
	if err := db.DB.Create(&player).Error; err != nil {
		return player, err
	}
	return player, nil
}

func (pc *PlayerController) GetPlayerOnline() (players []models.Players) {
	//TODO: corregir maun del futuro
	var playerOn []models.PlayersOnline
	db.DB.Find(&playerOn)
	for _, value := range playerOn {
		var playerTarget models.Players
		db.DB.Where("id = ?", value.PlayerID).First(&playerTarget)
		players = append(players, playerTarget)
	}
	return
}
