package media

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB)(*Repository){
	return &Repository{
		db: db,
	}
}

func (repository *Repository) CreateMedia(media *Media)(bool, error){
	err := repository.db.Create(media).Error
	if err == nil{
		return true, nil
	}
	return false, err
}