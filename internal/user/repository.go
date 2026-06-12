package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *Repository{
	return &Repository{
		db: db,
	}
}
func (r *Repository) Create(user *User, profile * Profile) error{
	return r.db.Transaction(
		func(tx *gorm.DB) error {
			if err:= tx.Create(user).Error; err != nil{
				return err
			}
			profile.UserID = user.ID
			return tx.Create(profile).Error
		})
}

func (r *Repository) FindByPhone(phone string)( *User, error){
	var user User

	if err:= r.db.Where("phone_number= ?", phone).First(&user).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindByEmail(email string) (*User, error){
	var user User

	//gán user tìm được vào &user, nếu có error->trả về nill, error
	err := r.db.Where("email= ?", email).First(&user).Error
	if err != nil{
		return nil, err
	}

	return &user, nil
}


func (r *Repository) FindByID(id uint)(*User, error){
	var user User

	err := r.db.Joins("Profile").Where("users.id=?", id).First(&user).Error

	if err != nil{
		return nil, err
	}

	return &user, nil
}