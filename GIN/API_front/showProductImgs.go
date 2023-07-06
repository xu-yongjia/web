package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// ShowProductImgs 商品详情图片接口
func ShowProductImgs(c *gin.Context) {
	service := ShowImgsService{}
	// 	res := service.Show(c.Param("id"))
	// 	c.JSON(200, res)
	// }
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(201, ERRRESPONSE(err.Error(), 201))

	}
}

// ShowImgsService 商品图片详情的服务
type ShowImgsService struct {
	ProductID int `form:"productID" json:"productID"`
}

// Show 商品图片
func (service *ShowImgsService) Show() Response {
	var imgs []DBstruct.ProductImg
	code := e.SUCCESS

	err := DBstruct.DB.Where("product_id=?", service.ProductID).Find(&imgs).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if BuildImgs(imgs) != nil {
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   BuildImgs(imgs),
		}
	} else {
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "",
		}
	}

}

// ProductImg 商品图片序列化器
type ProductImg struct {
	ID        uint   `json:"id"`
	ProductID int    `json:"product_id"`
	ImgPath   string `json:"img_path"`
	CreatedAt int64  `json:"created_at"`
}

// BuildImg 序列化商品图片
func BuildImg(item DBstruct.ProductImg) ProductImg {
	return ProductImg{
		ID:        item.ID,
		ProductID: item.ProductID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildImgs 序列化商品图片列表
func BuildImgs(items []DBstruct.ProductImg) (imgs []ProductImg) {
	for _, item := range items {
		img := BuildImg(item)
		imgs = append(imgs, img)
	}
	return imgs
}
