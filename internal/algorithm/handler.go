package algorithm

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) AlgorithmList(cntx *gin.Context){
	var filter AlgoFilter
	if err:= cntx.ShouldBindQuery(&filter); err != nil{
		cntx.JSON(400, gin.H{
			"message": "Filter không hợp lệ!",
		})
		return
	}
	algorithms, err:=handler.service.GetAlgorithms(&filter)
	if err != nil{
		cntx.JSON(500, gin.H{
			"message": "Lỗi truy vấn server!",
		})
		return
	}
	cntx.JSON(200, algorithms)
}

func (handler *Handler) GetAlgorithm (cntx *gin.Context){
	id, ok := cntx.Params.Get("id")
	if !ok {
		cntx.JSON(400, gin.H{"message": "Không tìm thấy id!"})
		return
	}
	id1, err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		cntx.JSON(400, gin.H{"message": "id không hợp lệ!"})
		return
	}

	algo, err := handler.service.repository.GetAlgorithm(uint(id1))

	if err != nil{
		cntx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	cntx.JSON(200, algo)
	
}