package algorithm

import (
	"github.com/HoanNghi16/Devall_backend/internal/pkg"
)

type AlgoFilter struct {
	Cursor uint   `form:"cursor"`
	Level  string `form:"level" binding:"omitempty,oneof=easy medium hard advanced"`
	Tags   []uint `form:"tags"`
}

type TagResponse struct {
	Tag
	Value int `json:"value"`
}

type AlgoResponse struct {
	pkg.BaseModel
	Name  		string `json:"name"`
	Level 		string `json:"level"` // easy, medium, hard, advanced	
	Description string `json:"description"`
	IsPublished bool   `json:"is_published"`
	IsSolved	bool   `json:"is_solved"`
	SolverID    *uint   `json:"solver_id"`
}

type RankingResponse struct{
	Rank        int `json:"rank"`
	RankerName  string `json:"ranker_name"`
	SolvedCount int `json:"solved_count"`
	EXP         uint64 `json:"exp"` 
}