package models

import "time"

type Session struct {
	ID        int        `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Token     []byte     `gorm:"column:token;size:16;not null"`       // BINARY(16)
	AccountID int        `gorm:"column:account_id;type:int;not null"` // INT
	IP        []byte     `gorm:"column:ip;size:16;not null"`          // VARBINARY(16)
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`    // TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	ExpiredAt *time.Time `gorm:"column:expired_at"`                   // TIMESTAMP NULL
}

// TableName overrides the default table name
func (Session) TableName() string {
	return "sessions"
}
