package models

type PlayersOnline struct {
	PlayerID uint `gorm:"column:player_id"`
}

func (PlayersOnline) TableName() string {
	return "players_online"
}
