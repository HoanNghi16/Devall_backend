package course

import (
	"errors"
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



	// ID               uint  `json:"id"`
	// Name             string`json:"name"`
	// Avatar           string`json:"avatar"`
	// Author           ResponseAuthor `json:"author"`
	// ShortDescription string `json:"short_description"`
	// CreatedAt        time.Time `json:"created_at"`
	// UpdatedAt        time.Time `json:"updated_at"`
	

	// //Course_user
	// Progress  		float32 `gorm:"not null; check: progress >= 0 and progress <= 1" json:"progress"`
	// DeletedAt 		*time.Time `gorm:"null" json:"deleted_at"`
	// IsActive  		bool `gorm:"default:true" json:"is_active"` //Dùng để hiển thị trong trang lịch sử hoặc ko
	// IsMarked  		bool `gorm:"default:false" json:"is_marked"`
	// LastAccessAt	time.Time `gorm:"not null; default:now()" json:"last_access_at"`



func (repository *Repository)FindAll(userID uint,cursor uint, topicIDs []uint, level string )([]ResponseCourse, error){
	var courses []ResponseCourse
	query := repository.db.Model(&Course{}).
			Select(`courses.id, courses.name,
					courses.avatar, courses.short_description, 
					cu.progress, cu.deleted_at, cu.is_active, 
					cu.is_marked, cu.last_access_at`). //Những fields của ResponseCourse
			Joins("Author").Where("courses.id > ? and courses.is_published = true", cursor)
	if len(topicIDs) > 0{
		query = query.Joins("join topic_courses tc on tc.course_id = courses.id").Where("tc.topic_id in ?", topicIDs).Distinct("courses.*")
	}

	if level != "" && level != "all"{
		query = query.Where("level = ?", level)
	}
	
	query = query.Joins("LEFT JOIN course_users cu on cu.course_id = courses.id", "user_id = ?", userID)


	err := query.Scan(&courses).Error

	if err != nil{
		return nil, err
	}

	return courses, nil
}


// Khoan sửa
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

func (repostiory *Repository) SelectHistories(userID uint, cursor uint)([]ResponseCourse, error){
	var courses []ResponseCourse

	query := repostiory.db.Model(&Course{}).
			Select(`courses.id, courses.name,
					courses.avatar, courses.short_description, 
					cu.progress, cu.deleted_at, cu.is_active, 
					cu.is_marked, cu.last_access_at`).Joins("Author").Joins("JOIN course_users cu on cu.course_id = courses.id", "user_id = ?", userID).Order("last_access_at DESC")

	var tempCourse *CourseUser
	if cursor != 0{
		if errTemp:= repostiory.db.Where("course_id = ? AND user_id =  ?", cursor, userID).Find(&tempCourse).Error; errTemp != nil{
			if errors.Is(errTemp, gorm.ErrRecordNotFound){
				return nil, errors.New("ID không hợp lệ!") 
			}
			return nil, errors.New("Kết nối server thất bại!")
		}
		query = query.Where("last_access_at > ", tempCourse.LastAccessAt)
	}

	if err := query.Limit(15).Scan(&courses).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return []ResponseCourse{}, nil
		}
		return nil, err
	}
	return courses, nil
}


func (repository *Repository) GetBookmarks(userID uint, cursor uint)([]ResponseCourse, error){
	var courses []ResponseCourse

	query := repository.db.Model(&Course{}).
			Select(`courses.id, courses.name,
					courses.avatar, courses.short_description, 
					cu.progress, cu.deleted_at, cu.is_active, 
					cu.is_marked, cu.last_access_at`).
			Joins("JOIN course_users cu on courses.id = cu.course_id", "user_id = ?", userID)
	if cursor != 0 {
		query = query.Where("id > ?", cursor)
	}
	
	if err := query.Scan(&courses).Error; err!=nil{
		return nil, err
	}
	return courses, nil
}