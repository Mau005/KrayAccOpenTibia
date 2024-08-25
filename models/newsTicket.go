package models

import "gorm.io/gorm"

type NewsTicket struct {
	gorm.Model
	IconID    uint8
	Ticket    string
	PlayersID int
	Player    Players `gorm:"foreignKey:ID;references:PlayersID"`
}
