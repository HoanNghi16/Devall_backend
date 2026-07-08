package user

import (
	"errors"

	"github.com/HoanNghi16/Devall_backend/internal/auth"
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


func (s *Service) Login(input *LoginRequest) (*TokenResponse, error) {
	user, _ := s.repository.FindByEmail(input.Email)

	if user == nil{
		user, _ = s.repository.FindByPhone(input.Phone)
		if (user == nil){
			return nil,errors.New("Email và số điện thoại không tồn tại!")
		}
		
	}
	if err:= bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(input.Password),
	); err != nil{
		return nil,errors.New("Sai mật khẩu")
	}

	access,acc_err := auth.GenerateAccess(user.ID, user.Role)
	refresh,refr_err := auth.GenerateRefresh(user.ID)
	if acc_err != nil || refr_err != nil{
		return nil, errors.New("Tạo token thất bại")
	}
	return &TokenResponse{
		Access: access,
		Refresh: refresh,
	}, nil
}


func (s *Service) GetProfile(id uint)(*UserResponse,error){
	user, err := s.repository.FindByID(id)

	if err != nil{
		return nil,err
	}

	profile := UserResponse{
		Email: user.Email,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		LastLogin: user.LastLogin,
		Profile: user.Profile,
	}

	return &profile, nil

}