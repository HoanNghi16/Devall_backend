package language

import "github.com/HoanNghi16/Devall_backend/internal/pkg"

type SupportedLanguage struct {
	pkg.BaseModel
	Name string `gorm:"not null"`
	Version string `gorm:"not null"`
	Description string `gorm:"not null"`
}