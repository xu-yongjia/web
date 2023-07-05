package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

func ListRanking(c *gin.Context) {
	service := ListRankingService{}
	res := service.List()
	c.JSON(200, res)
}

// ListRankingService 展示排行的服务
type ListRankingService struct {
}

// List 排行榜列表
func (service *ListRankingService) List() Response {
	products := []DBstruct.Product{}
	code := e.SUCCESS
	// total := 0
	if err := DBstruct.DB.Order("score desc").Limit(4).Find(&products).Error; err != nil {
		code = e.ERROR_DATABASE
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return BuildListResponse(BuildProducts(products), uint(0))
}
