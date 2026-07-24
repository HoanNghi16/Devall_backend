package course

import (
	"errors"
	"log"
	"time"

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

	err := query.Where("is_published = true").First(&course, id).Error

	now := time.Now()

	if err != nil{
		return nil, errors.New("Không tìm thấy khóa học!")
	}

	if (userID != 0){
		var courseUser CourseUser
		err := repository.db.Where("course_id = ? and user_id = ?", course.ID, userID).First(&courseUser).Error

		if errors.Is(err, gorm.ErrRecordNotFound){
			courseUser.CourseID = course.ID
			courseUser.UserID = userID
			courseUser.LastAccessAt = now
			repository.db.Create(&courseUser) // Nếu chưa thấy thì create
		}else{
			courseUser.LastAccessAt = now
			repository.db.Model(&courseUser).UpdateColumn("last_access_at", now) // Nếu thấy courseUser => Người dùng từng truy cập => udpdate last_access_at 
		}
		course.CourseUsers = []CourseUser{courseUser}
	}

	return &course, nil
}




func (repository *Repository)FindAll(userID uint,cursor uint, topicIDs []uint, level string )([]Course, error){
	var courses []Course
	query := repository.db.Joins("Author").Where("courses.id > ? and courses.is_published = true", cursor)
	if len(topicIDs) > 0{
		query = query.Joins("join topic_courses tc on tc.course_id = courses.id").Where("tc.topic_id in ?", topicIDs).Distinct("courses.*")
	}

	if level != "" && level != "all"{
		query = query.Where("level = ?", level)
	}

	if userID != 0{
		query = query.Preload("CourseUsers", "user_id = ?", userID)
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
func (repository *Repository) UpdateCourseUser(coureUser *CourseUser, columns []string)(bool){
	err := repository.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "course_id"},
			{Name: "user_id"},
		},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Create(coureUser).Error

	if err != nil{
		return false
	}

	return true
}	

func (repostiory *Repository) SelectHistories(userID uint)([]CourseUser, error){
	var courseUsers []CourseUser
	if err := repostiory.db.Where("user_id = ? and is_active = true", userID).Order("last_access_at DESC").Preload("Course").Preload("Course.Author").Find(&courseUsers).Error; err != nil{
		log.Print(err.Error())
		return nil, errors.New("Lỗi truy vấn dữ liệu!")
	}
	return courseUsers, nil

}