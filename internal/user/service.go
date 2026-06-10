package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Register(input *RegisterRequest) (error) {
	// Hàm này thành công nếu trả về User, nil
	// -> Nếu error == nil -> có user đã dùng email này
	_, err := s.repository.FindByEmail(input.Email) 

	if err == nil {
		return errors.New("Đã tồn tại Email!")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err!= nil{
		return err //Hash thất bại
	}

	user := User{
		Email: input.Email,
		Password: string(hash),
	}
	
	profile := Profile{
		Name: input.Name,
		PhoneNumber: input.Phone,
	}

	err = s.repository.Create(&user, &profile)
	if err!= nil{
		return err
	}


	return nil
}


func (s *Service) Login(input *LoginRequest) error {
	user, _ := s.repository.FindByEmail(input.Email)

	if user == nil{
		return errors.New("Email không tồn tại!")
	}

	if err:= bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(input.Password),
	); err != nil{
		return errors.New("Sai mật khẩu")
	}

	return nil
}