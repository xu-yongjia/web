package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// ListCategories 分类列表接口
func ListCategories(c *gin.Context) {
	service := ListCategoriesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(201, ERRRESPONSE(err.Error(), 201))
		// logging.Info(err)
	}
}

// ListCategoriesService 分类列表服务
type ListCategoriesService struct {
	Limit     int  `form:"limit" json:"limit"`
	Start     int  `form:"start" json:"start"`
	CanteenID uint `form:"placeID" json:"placeID"`
}

// List 分类列表
func (service *ListCategoriesService) List() Response {

	categories := []DBstruct.Category{}

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	//fmt.Printf("service.CanteenID: %v\n", service.CanteenID)
	if service.CanteenID == 0 { //返回全部
		if err := DBstruct.DB.Model(DBstruct.Category{}).Count(&total).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Limit(service.Limit).Offset(service.Start).Find(&categories).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := DBstruct.DB.Model(DBstruct.Category{}).Where("canteen_id=?", service.CanteenID).Count(&total).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Where("canteen_id=?", service.CanteenID).Limit(service.Limit).Offset(service.Start).Find(&categories).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return BuildListResponse(BuildCategories(categories), uint(total))

}

// Category 分类序列化器
type Category struct {
	ID           uint   `json:"id"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

// BuildCategory 序列化分类
func BuildCategory(item DBstruct.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

// BuildCategories 序列化分类列表
func BuildCategories(items []DBstruct.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}

// TrackedErrorResponse 有追踪信息的错误响应
