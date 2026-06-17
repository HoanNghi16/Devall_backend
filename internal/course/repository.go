package course

import (
	"errors"
	"log"

	"github.com/HoanNghi16/Devall_backend/internal/user"
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
	log.Println(id)
	err := repository.db.Preload("Lessons").Preload("Lessons.ContentBlocks").Find(&course, id).Error
	if err != nil{
		return nil, err
	}
	if course.ID == 0{
		return nil, errors.New("Không tìm thấy khóa học!")
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

func (repository *Repository)GetMyCourses(userID uint)([]Course, error){
	var courses []Course
	var user *user.User
	err1 := repository.db.Joins("Profile").Where("users.id = ?", userID).Find(&user).Error
	if err1 == nil{
		err := repository.db.Where("author_id = ?", user.Profile.ID).Find(&courses).Error
		if err != nil{
			return nil, err
		}
		return courses,nil
	}
	return nil, err1
}