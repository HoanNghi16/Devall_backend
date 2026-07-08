package media

import "mime/multipart"

type MediaRequest struct {
	Name string `form:"name" binding:"omitempty,max=100"`
	File *multipart.FileHeader `form:"file" binding:"required"`
	Type string `form:"type" binding:"oneof=image video"`
}