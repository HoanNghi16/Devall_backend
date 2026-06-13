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

//Lấy chi tiết khóa học
//Dùng Preload để lấy Danh sách bài học trước
//Dùng Preload để lấy ContentBlocks trong Lessons
//->Find để đưa vào course
func (repository *Repository)GetCourse(id uint)(*Course, error){
	var course Course
	err := repository.db.Preload("Lessons").Preload("Lessons.ContentBlocks").Find(&course, id).Error
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