package models

type PlayerItem struct {
	PlayerID   int    `gorm:"column:player_id;primaryKey"`
	PID        int    `gorm:"column:pid;primaryKey"`
	SID        int    `gorm:"column:sid;primaryKey"`
	ItemType   uint16 `gorm:"column:itemtype"`   // UNSIGNED SMALLINT
	Count      int16  `gorm:"column:count"`      // SMALLINT (puede ser negativo en teoría)
	Attributes []byte `gorm:"column:attributes"` // BLOB
}

// TableName overrides the default table name
func (PlayerItem) TableName() string {
	return "player_items"
}
