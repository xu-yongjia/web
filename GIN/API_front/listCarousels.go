package api1

import (
	"fmt"
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

// ListCarouselsService 轮播图服务
type ListCarouselsService struct {
}

func (service *ListCarouselsService) List() Response {
	Canteens := []DBstruct.Canteen{}
	code := e.SUCCESS

	if err := DBstruct.DB.Find(&Canteens).Error; err != nil {
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
		Data:   BuildCarousels(Canteens),
	}
}

// Carousel 轮播图序列化器
type Carousel struct {
	ImgPath     string `json:"img_path"`
	CanteenID   uint   `json:"canteen_id"`
	CanteenName string `json:"canteen_name"`
}

// BuildCarousel 序列化轮播图
func BuildCarousel(item DBstruct.Canteen) Carousel {
	return Carousel{
		ImgPath:     item.ImgPath,
		CanteenID:   uint(item.CanteenID),
		CanteenName: item.CanteenName,
	}
}

// BuildCarousels 序列化轮播图列表
type JsonCarousel struct {
	Carousels []Carousel `json:"carousel"`
	Board     []string   `json:"board"`
}

func BuildCarousels(items []DBstruct.Canteen) (jc JsonCarousel) {
	var carousels []Carousel
	var msg []string
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
		msg = append(msg, item.Board)
	}
	fmt.Print(msg)
	jc.Board = msg
	jc.Carousels = carousels
	return
}
