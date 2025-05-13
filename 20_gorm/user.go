package gorm

import (
	"time"

	"gorm.io/gorm"
)

// Nama struct dan semua fieldnya sudah sesuai konvensi gorm,
// jadi sebenarnya TableName() dan semua tag gorm: tidak diperlukan,
// autoCreateTime dan autoUpdateTime saja secara default sudah :true.
/*
	type User struct {
		ID          string    `gorm:"primary_key;column:id;"`
		Password    string    `gorm:"column:password"`
		Name        string    `gorm:"column:name"`
		CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;"`
		UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
	}

	func (u *User) TableName() string {
		return "users"
	}
*/

type User struct {
	ID            string    `gorm:"primary_key;column:id;<-:create"`
	Password      string    `gorm:"column:password"`
	Name          Name      `gorm:"embedded"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
	Information   string    `gorm:"-"`
	Wallet        Wallet    `gorm:"foreignKey:user_id;references:id"`
	Addresses     []Address `gorm:"foreignKey:user_id;references:id"`
	LikedProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		u.ID = "user-" + time.Now().Format("20060102150405")
	}
	return nil
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

type UserLog struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	UserID    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int    `gorm:"column:created_at;autoCreateTime:milli;<-:create"`
	UpdatedAt int    `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *UserLog) TableName() string {
	return "user_logs"
}
