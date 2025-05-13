package gorm

import "time"

type Address struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    string    `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
}

func (Address) TableName() string {
	return "addresses"
}
