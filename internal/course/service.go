package course

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