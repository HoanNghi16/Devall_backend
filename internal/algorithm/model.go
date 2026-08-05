package algorithm

import (
	"github.com/HoanNghi16/Devall_backend/internal/language"
	"github.com/HoanNghi16/Devall_backend/internal/pkg"
	"github.com/HoanNghi16/Devall_backend/internal/user"
)

type Algorithm struct {
	pkg.BaseModel
	Name  		string `gorm:"not null;unique" json:"name"`
	Level 		string `gorm:"not null" json:"level"` // easy, medium, hard, advanced	
	Description string `gorm:"not null" json:"description"`
	Tags		[]Tag  `gorm:"many2many:algo_tags" json:"tags"`
	IsPublished bool   `gorm:"not null;default:false" json:"is_published"`
	SolvingHistories []SolvingHistory `gorm:"foreignKey:AlgorithmID" json:"solving_histories"`
}

type SolvingHistory struct {
	pkg.BaseModel
	IsSolved 	bool		`gorm:"not null;default:false" json:"is_solved"`
	AlgorithmID uint		`gorm:"not null"`
	Algorithm 	Algorithm 	`gorm:"foreignKey:AlgorithmID"`
	SolverID   	uint		`gorm:"not null"`
	Solver 		user.User 	`gorm:"foreignKey:SolverID"`
	Script      string		`gorm:"not null"`
	Runtime     string		`gorm:"not null"`
	LanguageID 	uint 		`gorm:"not null"`
	Language 	language.SupportedLanguage `gorm:"foreignKey:LanguageID"`
}

type Tag struct {
	ID 	 	uint	`json:"id"`
	Name 	string  `json:"name"`
}