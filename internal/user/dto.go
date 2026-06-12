package user

import "time"

type TokenResponse struct{
	Access string
	Refresh string
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"required,min=10,numeric"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"string"`
	Password string `json:"password" binding:"min=8"`
}

type ProfileResponse struct {
	ID        uint   
	Email     string 
	Role      string 
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin time.Time
	Profile   Profile
}