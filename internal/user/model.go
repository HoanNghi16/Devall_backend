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
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint `gorm:"unique;not null" json:"user_id"`	
	EXP uint64 `gorm:"default:0" json:"exp"`
	Avatar string `gorm:"default:''" json:"avatar"`
	Name string `gorm:"default:''" json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
}