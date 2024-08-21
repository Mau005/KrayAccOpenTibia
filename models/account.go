package models

type Account struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"type:varchar(32)"`
	Password      string
	Secret        string
	Type          uint
	PremiumEndsAt int `gorm:"column:premium_ends_at"`
	Email         string
	Creation      int
	PremiumDays   int64 `gorm:"column:premiumdays"`
	LastDay       int64 `gorm:"column:lastday"`
}
