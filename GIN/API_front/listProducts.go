package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// TrackedErrorResponse 有追踪信息的错误响应

// ListProducts 商品列表接口
func ListProducts(c *gin.Context) {
	var service ListProductsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(201, ERRRESPONSE(err.Error(), 201))
	}

}

// ListProductsService 商品列表服务
type ListProductsService struct {
	Limit int `form:"limit" json:"limit"` //分页查找
	Start int `form:"start" json:"start"`
	// CanteenID  uint `form:"canteen_id" json:"canteen_id"`
	CanteenID uint `form:"placeID" json:"placeID"`
	// CategoryID uint `form:"category_id" json:"category_id"`
	CategoryID uint `form:"categoryID" json:"categoryID"`
}

// List 商品列表
func (service *ListProductsService) List() Response {
	products := []DBstruct.Product{}

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	if service.CanteenID == 0 && service.CategoryID == 0 { //没选定食堂和类型
		if err := DBstruct.DB.Model(DBstruct.Product{}).Count(&total).Error; err != nil {
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else if service.CategoryID == 0 && service.CanteenID != 0 { //选定食堂,没选定类型
		if err := DBstruct.DB.Model(DBstruct.Product{}).Where("canteen_id=?", service.CanteenID).Count(&total).Error; err != nil {
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Where("canteen_id=?", service.CanteenID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			//logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else if service.CategoryID != 0 && service.CanteenID == 0 { //没选食堂，选定类型
		if err := DBstruct.DB.Model(DBstruct.Product{}).Where("category_id=?", service.CategoryID).Count(&total).Error; err != nil {
			//logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Where("category_id=?", service.CategoryID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			//logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

	} else { //都选了
		if err := DBstruct.DB.Model(DBstruct.Product{}).Where("category_id=? AND canteen_id=?", service.CategoryID, service.CanteenID).Count(&total).Error; err != nil {
			//logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Where("category_id=? AND canteen_id=?", service.CategoryID, service.CanteenID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			//logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return BuildListResponse(BuildProducts(products), uint(total))
}
