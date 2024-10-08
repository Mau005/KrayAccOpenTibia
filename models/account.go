package models

type Account struct {
	ID            int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"column:name;unique;size:32" json:"name"`
	Password      string    `gorm:"column:password;size:40" json:"password"` //Check not send for password, use api internal
	Secret        *string   `gorm:"column:secret;size:16" json:"secret"`
	Type          int       `gorm:"column:type;default:1" json:"type"`
	PremiumEndsAt uint      `gorm:"column:premium_ends_at;default:0" json:"premium_ends_at"`
	Email         string    `gorm:"column:email;not null;unique;size:255;default:''" json:"email"`
	Creation      int       `gorm:"column:creation;default:0" json:"creation"`
	Players       []Players `gorm:"foreignKey:AccountID" json:"players"`
}
