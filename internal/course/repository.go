package course

import (
	"errors"
	"log"

	"github.com/HoanNghi16/Devall_backend/internal/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
func (repository *Repository)GetCourse(id uint, userID uint)(*Course, error){
	var course Course
	log.Println(userID)
	query := repository.db.Preload("Lessons").Preload("Lessons.ContentBlocks")

	if (userID != 0){
		query = query.Preload("CourseUsers", "user_id = ?", userID)
	}

	err := query.Where("is_published = true").Find(&course, id).Error

	if err != nil{
		return nil, err
	}
	if course.ID == 0{
		return nil, errors.New("Không tìm thấy khóa học!")
	}
	return &course, nil
}




func (repository *Repository)FindAll(cursor uint, topicIDs []uint, level string )([]Course, error){
	var courses []Course
	query := repository.db.Joins("Author").Where("courses.id > ? and courses.is_published = true", cursor)
	if len(topicIDs) > 0{
		query = query.Joins("join topic_courses tc on tc.course_id = courses.id").Where("tc.topic_id in ?", topicIDs).Distinct("courses.*")
	}

	if level != "" && level != "all"{
		query = query.Where("level = ?", level)
	}

	err := query.Find(&courses).Error

	if err != nil{
		return nil, err
	}

	return courses, nil
}

func (repository *Repository)GetMyCourses(userID uint)([]Course, error){
	var courses []Course
	var user *user.User
	err1 := repository.db.Joins("Profile").Where("users.id = ?", userID).First(&user).Error
	if err1 == nil{
		err := repository.db.Joins("Author").Where("author_id = ?", user.Profile.ID).Find(&courses).Error
		if err != nil{
			return nil, err
		}
		return courses,nil
	}
	return nil, err1
}


func (repository *Repository) CreateMyCourse(course *Course)(error){
	return repository.db.Create(course).Error
}


// Lấy danh sách Topic
func (repository *Repository) GetTopics ()([]Topic,error){
	var topics []Topic
	err:=repository.db.Find(&topics).Error
	return topics, err
}


// Cập nhật khóa học
func (repository *Repository) UpdateCourseUser(coureUser *CourseUser)(bool){
	err := repository.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "course_id"},
			{Name: "user_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"progress",
			"updated_at",
			"is_marked",
		}),
	}).Create(coureUser).Error

	if err != nil{
		return false
	}

	return true
}	