package algorithm

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) GetAlgorithms(filter *AlgoFilter) ([]Algorithm, error) {
	algorithms, err := service.repository.GetAlgos(filter.Tags, filter.Level, filter.Cursor)
	if err != nil {
		return nil, err
	}
	return algorithms, nil
}