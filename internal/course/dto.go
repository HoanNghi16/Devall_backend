package course

import (
	"time"
)

type ResponseAuthor struct{
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

type ResponseCourse struct { //Để json.Marshal() trả về đúng tên fields
	ID               uint  `json:"id"`
	Name             string`json:"name"`
	Avatar           string`json:"avatar"`
	Author           ResponseAuthor `json:"author"`
	ShortDescription string `json:"short_description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}