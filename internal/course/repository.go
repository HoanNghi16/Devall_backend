package course

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository){
	return &Repository{
		db: db,
	}
}

//Chưa chạy được!
func (repository *Repository)GetCourse(id uint)(*Course, error){
	var course Course
	err := repository.db.Joins("Lessons").Joins("ContentBlocks").Where("courses.id = ?", id).First(&course).Error
	if err != nil{
		return nil, err
	}
	return &course, nil
}

func (repository *Repository)FindAll(cursor uint)([]Course, error){
	var courses []Course
	err := repository.db.Joins("Author").Where("courses.id > ? and is_published = true", cursor).Limit(12).Find(&courses).Error
	if err != nil{
		return nil, err
	}
	return courses, nil
}