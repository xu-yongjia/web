package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// ShowProduct 商品详情接口
func ShowProduct(c *gin.Context) {
	var service ShowProductService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(201, ERRRESPONSE(err.Error(), 201))
		//logging.Info(err)
	}
}

// ShowProductService 商品详情的服务
type ShowProductService struct {
	ProductID string `form:"productID" json:"productID"`
}

// Show 商品

func (service *ShowProductService) Show() Response {
	var product DBstruct.Product
	code := e.SUCCESS

	// id, _ := strconv.Atoi(service.ProductID)
	if err := DBstruct.DB.First(&product, service.ProductID).Error; err != nil {
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
		Data:   BuildProduct(product),
	}
}
