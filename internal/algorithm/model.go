package algorithm

import (
	"github.com/HoanNghi16/Devall_backend/internal/user"
	"gorm.io/gorm"
)

type Algorithm struct {
	gorm.Model
	Name  		string `gorm:"not null;unique" json:"name"`
	Level 		string `gorm:"not null" json:"level"` // easy, medium, hard, advanced	
	Description string `gorm:"not null" json:"description"`
	Tags		[]Tag  `gorm:"many2many:algo_tags" json:"tags"`
	IsPublished bool   `gorm:"not null;default:false" json:"is_published"`
}

type SolvingHistory struct {
	gorm.Model
	AlgorithmID uint
	Algorithm 	Algorithm 	`gorm:"foreignKey:AlgorithmID"`
	SolverID   	uint		`gorm:"not null"`
	Solver 		user.User 	`gorm:"foreignKey:SolverID"`
	Script      string		`gorm:"not null"`
	Runtime     string		`gorm:"not null"`
}

type Tag struct {
	ID 	 	uint
	Name 	string 
}