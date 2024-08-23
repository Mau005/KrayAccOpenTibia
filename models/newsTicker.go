package models

import "gorm.io/gorm"

type NewsTicker struct {
	gorm.Model
	Icon        string
	Tickets     string
	ByCharacter string
}
