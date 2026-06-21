package course

import (
	"errors"

	"github.com/HoanNghi16/Devall_backend/internal/user"
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

func (service *Service) CourseFullService(id uint) (*Course, error) {
	course, err := service.repository.GetCourse(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (service *Service) MyCourseService(userID uint) ([]Course, error) {
	courses, err := service.repository.GetMyCourses(userID)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (service *Service) CreateMyCourse(userID uint, input *RequestCourse) error {
	userRepository := user.NewRepository(service.repository.db)

	user, err := userRepository.FindByID(userID) 
	if err != nil{
		return errors.New("ID người dùng không hợp lệ!")
	}
	course := input.ParseCourse()
	course.Author = user.Profile
	if err := service.repository.CreateMyCourse(&course); err != nil {
		return err
	}
	return nil
}