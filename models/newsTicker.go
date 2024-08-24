package models

import "gorm.io/gorm"

type NewsTicket struct {
	gorm.Model
	Icon     string
	IconID   uint8
	Ticket   string
	PlayerID uint
	Player   Player `gorm:"foreignKey:ID"`
}
