package user

import "time"

type TokenResponse struct{
	Access string
	Refresh string
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone_number" binding:"required,min=10,numeric"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone_number"`
	Password string `json:"password" binding:"min=8"`
}

type UserResponse struct {
	ID 		  uint		`json:"id"`
	Email     string 	`json:"email"`
	Role      string 	`json:"role"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	LastLogin time.Time	`json:"last_login"`
	Profile   Profile	`json:"profile"`
}