package media

import (
	"github.com/HoanNghi16/Devall_backend/internal/user"
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	PublicID string `gorm:"not null;unique" json:"public_id"`
	URL string `gorm:"not null" json:"url"`
	Type string `gorm:"not null" json:"media_type"`
	UploadedByID uint `gorm:"not null" json:"uploaded_by_id"`
	UploadedBy user.User `gorm:"foreignKey:UploadedByID" json:"uploaded_by"`
}