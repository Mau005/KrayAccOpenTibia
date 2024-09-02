package models

// player_id
// time
// level
// killed_by
// is_player
// mostdamage_by
// mostdamage_is_player
// unjustified
// mostdamage_unjustified
type PlayerDeaths struct {
	PlayersID              int     `gorm:"column:player_id;type:int(11)" json:"player_id"`
	Time                   int     `gorm:"column:time" json:"time"`
	Level                  int     `gorm:"column:level" json:"level"`
	KilledBy               string  `gorm:"column:killed_by" json:"killed_by"`
	IsPLayer               int     `gorm:"column:is_player" json:"is_player"`
	MostDamageBy           string  `gorm:"column:mostdamage_by" json:"mostdamage_by"`
	MostDamageIsPLayer     int     `gorm:"column:mostdamage_is_player" json:"mostdamage_is_player"`
	Unjustified            int     `gorm:"column:unjustified" json:"unjustified"`
	MonstDamageUnjustified int     `gorm:"column:mostdamage_unjustified" json:"mostdamage_unjustified"`
	Player                 Players `gorm:"foreignKey:ID;references:PlayersID"`
}

func (PlayerDeaths) TableName() string { return "player_deaths" }
