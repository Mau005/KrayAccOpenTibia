package models

import "gorm.io/gorm"

//use check unique player create
type PlayersNames struct {
	gorm.Model
	Name      string `gorm:"unique"`
	World     string
	AccountID int
}
