package algorithm

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB)(*Repository){
	return &Repository{
		db: db,
	}
}

func (repository *Repository) GetAlgos(tags []uint, level string,cursor uint)([]Algorithm, error){
	query_set := repository.db.Where("algorithms.id > ?", cursor).Where("is_published = true").Limit(12)
	if (len(tags) > 0){
		query_set = query_set.Joins("algo_tags").Where("Where algo_tags.tag_id in ?", tags).Distinct("Algorithm.*")
	}
	if level != ""{
		query_set = query_set.Where("level = ?", level)
	}
	var algorithms []Algorithm

	if err:= query_set.Find(&algorithms).Error; err != nil{
		return nil,err
	}
	return algorithms, nil
}

func (repository *Repository) GetAlgorithm(id uint)(*Algorithm, error){
	var algo Algorithm
	if err := repository.db.Preload("Tags").Where("is_published = true").Find(&algo, id).Error; err != nil{
		return nil, errors.New("Lỗi kết nối database!")
	}
	return &algo, nil
}