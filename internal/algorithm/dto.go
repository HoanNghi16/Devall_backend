package algorithm

type AlgoFilter struct {
	Cursor uint   `form:"cursor"`
	Level  string `form:"level" binding:"omitempty,oneof=easy medium hard advanced"`
	Tags   []uint `form:"tags"`
}