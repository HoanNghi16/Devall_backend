package user

import "time"

type User struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role string `gorm:"default:student"`
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin time.Time
	Profile Profile `gorm:"foreignKey:UserID"`
}

type Profile struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"unique;not null"`
	EXP uint64 `gorm:"default:0"`
	Avatar string `gorm:"default:''"`
	Name string `gorm:"default:''"`
	DateOfBirth time.Time
	PhoneNumber string
}