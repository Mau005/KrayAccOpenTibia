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
	PlayersID              int32   `gorm:"column:player_id;type:int(11)" json:"player_id"`
	Time                   int     `gorm:"column:time" json:"time"`
	Level                  int     `gorm:"column:level" json:"level"`
	KilledBy               string  `gorm:"column:killed_by;type:varchar(255)"`
	IsPLayer               int     `gorm:"column:is_player" json:"is_player"`
	MostDamageBy           string  `gorm:"column:mostdamage_by;type:varchar(100)"`
	MostDamageIsPLayer     int     `gorm:"column:mostdamage_is_player" json:"mostdamage_is_player"`
	Unjustified            int     `gorm:"column:unjustified" json:"unjustified"`
	MonstDamageUnjustified int     `gorm:"column:mostdamage_unjustified" json:"mostdamage_unjustified"`
	Player                 Players `gorm:"foreignKey:PlayersID;references:ID" json:"player"`
}

func (PlayerDeaths) TableName() string { return "player_deaths" }
