package course

import (
	"errors"

	"github.com/HoanNghi16/Devall_backend/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) ListCourseService(cursor uint, topicIDs []uint, level string) ([]ResponseCourse, error) {
	courses, err := service.repository.FindAll(cursor, topicIDs, level)
	if err != nil {
		return nil, err
	}
	var course *Course
	return course.ToResponseDataList(courses), nil
}


func (service *Service) CourseFullService(id uint, userID uint) (*Course, error) {
	course, err := service.repository.GetCourse(id, userID)
	if err != nil {
		return nil, err
	}
	course.Password = ""
	return course, nil
}



func (service *Service) MyCourseService(userID uint) ([]ResponseCourse, error) {
	courses, err := service.repository.GetMyCourses(userID)
	if err != nil {
		return nil, err
	}
	var course *Course
	return course.ToResponseDataList(courses), nil
}



func (service *Service) CreateMyCourse(userID uint, input *RequestCourse) error {
	userRepository := user.NewRepository(service.repository.db)

	user, err := userRepository.FindByID(userID) 
	if err != nil{
		return errors.New("ID người dùng không hợp lệ!")
	}

	hash, hashErr := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if hashErr != nil{
		return errors.New("Lỗi hệ thống xử lý mật khẩu!")
	}
	input.Password = string(hash)
	course := input.ParseCourse()
	course.Author = user.Profile
	if err := service.repository.CreateMyCourse(&course); err != nil {
		return err
	}
	return nil
}


func (service *Service) GetTopics ()([]Topic, error){
	topcics, err := service.repository.GetTopics()

	if err != nil{
		return nil, errors.New("Không tìm thấy dữ liệu")
	}
	return topcics, nil
}

func(service *Service) UpdateCoureUser(userID uint, input *RequestCourseUser)(error){
	
	courseUser, columns := input.ParseCourseUser()

	courseUser.UserID = userID

	ok := service.repository.UpdateCourseUser(&courseUser, columns)

	if ok{
		return nil
	}

	return errors.New("Thêm hoặc sửa dữ liệu thất bại!")
}

func(service *Service) GetHistories(userID uint)([]ResponseCourse, error){
	courseUsers, err := service.repository.SelectHistories(userID)
	if err != nil{
		return nil,err
	}

	responseCourses := make([]ResponseCourse, len(courseUsers))

	for index, courseUser := range courseUsers{
		responseCourses[index] = ResponseCourse{
			ID: courseUser.CourseID ,
			Name: courseUser.Course.Name,
			Avatar: courseUser.Course.Avatar,
			Author: ResponseAuthor{Name: courseUser.Course.Author.Name, Avatar: courseUser.Course.Author.Avatar},
			ShortDescription: courseUser.Course.ShortDescription,
			CourseUser: &courseUser,
		}
	}

	return responseCourses, nil
}