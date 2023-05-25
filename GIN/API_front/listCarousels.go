package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// ListCarouselsService 视频列表服务
type ListCarouselsService struct {
}

func ListCarousels(c *gin.Context) {
	service := ListCarouselsService{}
	if err := c.ShouldBind(&service); err == nil { //获取参数（并没有参数）
		res := service.List() //处理请求获得返回的结构体
		c.JSON(200, res)
	} else {
		c.JSON(200, ERRRESPONSE(err.Error(), 201))
	}
}

// List 视频列表
func (service *ListCarouselsService) List() EResponse {
	carousels := []DBstruct.Carousel{} //容纳轮播图列表的数组
	code := e.SUCCESS                  //返回的状态码

	if err := DBstruct.DB.Find(&carousels).Error; err != nil { //将数据库中所有轮播图都添加到数组中
		code = e.ERROR_DATABASE //查找出错，返回数据库错误的状态码
		return EResponse{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return EResponse{ //查找成功，返回查找结果
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   BuildCarousels(carousels), //将查找结果序列化
	}
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	EResponse
	TrackID string `json:"track_id"`
}

// Carousel 轮播图序列化器
type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductID uint   `json:"product_id"`
	CreatedAt int64  `json:"created_at"`
}

// BuildCarousel 序列化轮播图
func BuildCarousel(item DBstruct.Carousel) Carousel { //将DBstruct.Carousel转化为具有json字段名的Carousel
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		ProductID: item.ProductId,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildCarousels 序列化轮播图列表
func BuildCarousels(items []DBstruct.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(item)         //查询结果，将其中每个结构体都转化为一个序列化的结构体
		carousels = append(carousels, carousel) //添加到另一个数组中
	}
	return carousels
}
