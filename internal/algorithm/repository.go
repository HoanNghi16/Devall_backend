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

// type AlgoResponse struct {
// 	pkg.BaseModel
// 	Name  		string `json:"name"`
// 	Level 		string `json:"level"` // easy, medium, hard, advanced	
// 	Description string `json:"description"`
// 	Tags		[]Tag  `json:"tags"`
// 	IsPublished bool   `json:"is_published"`
// 	IsSolved	bool   `json:"is_solved"`
// 	SolverID    uint   `json:"solver_id"`
// }


func (repository *Repository) GetAlgos(userID uint, tags []uint, level string,cursor uint)([]AlgoResponse, error){
	query_set := repository.db.Model(&Algorithm{}).Select("algorithms.*, CASE WHEN sh.is_solved IS NOT NULL THEN sh.is_solved ELSE false END as is_solved,  sh.solver_id").Where("algorithms.id > ? AND is_published = true", cursor).Limit(12)
	if (len(tags) > 0){
		query_set = query_set.Where("algorithms.id in (SELECT DISTINCT algorithm_id FROM algo_tags WHERE tag_id IN ?)", tags)
	}
	if level != ""{
		query_set = query_set.Where("level = ?", level)
	}
	var algorithms []AlgoResponse

	if userID != 0{
		query_set = query_set.Joins("left join solving_histories sh on sh.algorithm_id = algorithms.id", "sh.solver_id = ? and is_solved = true", userID)
	}else{
		query_set = query_set.Joins("left join (select * from solving_histories where solver_id = 0) sh on sh.algorithm_id = algorithms.id")
	}
	
	if err:= query_set.Scan(&algorithms).Error; err != nil{
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

func (repository *Repository) GetTagsWithAlgo()([]TagResponse, error){
	query_set := repository.db.Model(&Tag{}).Select("tags.*, count(algo_tags.*) as value").Joins("left join algo_tags on algo_tags.tag_id = tags.id").Group("tags.id")

	var tags []TagResponse

	if err := query_set.Scan(&tags).Error; err != nil{
		return nil, err
	}

	return tags, nil
}