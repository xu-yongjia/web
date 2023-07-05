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
	if service.CanteenID == 0 { //没有限制食堂，则返回全部
		if err := DBstruct.DB.Model(DBstruct.Category{}).Group("category_id").Count(&total).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Limit(service.Limit).Offset(service.Start).Order("category_id desc").Find(&categories).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else { //如果限制了食堂，那么类别不会重复
		if err := DBstruct.DB.Model(DBstruct.Category{}).Where("canteen_id=?", service.CanteenID).Count(&total).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Where("canteen_id=?", service.CanteenID).Limit(service.Limit).Offset(service.Start).Order("category_id desc").Find(&categories).Error; err != nil {
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
}

// BuildCategory 序列化分类
func BuildCategory(item DBstruct.Category) Category {
	return Category{
		ID:           item.CategoryID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
	}
}

// BuildCategories 序列化分类列表
func BuildCategories(items []DBstruct.Category) (categories []Category) {
	var temp uint = 0
	for _, item := range items {
		if temp != item.CategoryID { //跳过重复categoryId
			temp = item.CategoryID
			category := BuildCategory(item)
			categories = append(categories, category)
		}
	}
	return categories
}

// TrackedErrorResponse 有追踪信息的错误响应
