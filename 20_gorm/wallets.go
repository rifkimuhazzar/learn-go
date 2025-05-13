package gorm

import "time"

type Wallet struct {
	ID        string    `gorm:"column:id;primaryKey"`
	UserID    string    `gorm:"column:user_id"`
	Balance   uint64    `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      *User     `gorm:"foreignKey:user_id;references:id"`
}

func (Wallet) TableName() string {
	return "wallets"
}
