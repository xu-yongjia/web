package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// ListCarousels 轮播图列表接口
func ListCarousels(c *gin.Context) {
	service := ListCarouselsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(201, ERRRESPONSE(err.Error(), 201))
	}
}

// ListCarouselsService 视频列表服务
type ListCarouselsService struct {
}

// List 视频列表
func (service *ListCarouselsService) List() Response {
	carousels := []DBstruct.Carousel{}
	code := e.SUCCESS

	if err := DBstruct.DB.Find(&carousels).Error; err != nil {
		code = e.ERROR_DATABASE
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   BuildCarousels(carousels),
	}
}

// Carousel 轮播图序列化器
type Carousel struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	CanteenID uint   `json:"product_id"` //json名是历史遗留问题，不要在意
	CreatedAt int64  `json:"created_at"`
}

// BuildCarousel 序列化轮播图
func BuildCarousel(item DBstruct.Carousel) Carousel {
	return Carousel{
		ID:        item.ID,
		ImgPath:   item.ImgPath,
		CanteenID: item.CanteenId,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildCarousels 序列化轮播图列表
func BuildCarousels(items []DBstruct.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
