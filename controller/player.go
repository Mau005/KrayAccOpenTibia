package controller

import (
	"fmt"

	"github.com/Mau005/KrayAccOpenTibia/db"
	"github.com/Mau005/KrayAccOpenTibia/models"
	"github.com/Mau005/KrayAccOpenTibia/utils"
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
func (pc *PlayerController) GetPlayerID(id int) (player models.Players) {
	db.DB.Where("id = ?", id).First(&player)
	return
}
func (pc *PlayerController) GetAllPlayer() (player []models.Players) {
	db.DB.Find(&player)
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

func (pc *PlayerController) IndexHighScore(indexHighScore int) (response string) {
	switch indexHighScore {
	case utils.FirstHighScore:
		return "skill_fist"
	case utils.ClubHighScore:
		return "skill_club"
	case utils.SwordHighScore:
		return "skill_sword"
	case utils.AxeHighScore:
		return "skill_axe"
	case utils.DistHighScore:
		return "skill_dist"
	case utils.ShieldHighScore:
		return "skill_shielding"
	case utils.FishingHighScore:
		return "skill_fishing"
	case utils.MagLevelHighScore:
		return "maglevel"
	default:
		return "level"
	}
}

func (pc *PlayerController) GetHighScore(indexHighScore int) (players []models.Players) {
	target := pc.IndexHighScore(indexHighScore)
	db.DB.Limit(utils.LimitRecordHighScore).Order(fmt.Sprintf("%s desc", target)).Find(&players)
	return
}

func (pc *PlayerController) GetNameWorld(name string) models.PlayersNames {
	var playerName models.PlayersNames
	db.DB.Where("name = ?", name).First(&playerName)
	return playerName
}

func (pc *PlayerController) GetPlayerDeath() (deaths []models.PlayerDeaths) {
	db.DB.Preload("Player").Order("time desc").Limit(50).Find(&deaths)
	return
}
