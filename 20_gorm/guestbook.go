package gorm

import "time"

type GuestBook struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Message   string    `gorm:"column:message"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (GuestBook) TableName() string {
	return "guest_books"
}
