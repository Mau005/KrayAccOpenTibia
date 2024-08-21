package models

type Player struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(255)"`
	Level      uint16
	Sex        int
	Vocation   uint8
	LookType   uint16 `gorm:"column:looktype"`
	LookHead   uint16 `gorm:"column:lookhead"`
	LookBody   uint16 `gorm:"column:lookbody"`
	LookLegs   uint16 `gorm:"column:looklegs"`
	LookFeet   uint16 `gorm:"column:lookfeet"`
	LookAddons uint16 `gorm:"column:lookaddons"`
	Deleted    bool
	LastLogin  int64 `gorm:"column:lastlogin"`
}
