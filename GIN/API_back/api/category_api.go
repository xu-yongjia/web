package api

import (
	api2 "gintest/API_back"
	"gintest/API_back/vo"
	"gintest/DBstruct"
	"github.com/gin-gonic/gin"
)

type listParams struct {
	CanteenId *int `json:"canteen_id" binding:"required"`
}

type addCategoryParams struct {
	CanteenId    uint   `json:"canteen_id" binding:"required"`
	CategoryName string `json:"category_name" binding:"required"`
}

func (p *addCategoryParams) mergeFieldValue(model *DBstruct.Category) {
	model.CanteenID = p.CanteenId
	model.CategoryName = p.CategoryName
}

type updateCategoryParams struct {
	CategoryId   uint   `json:"category_id" binding:"required"`
	CanteenId    uint   `json:"canteen_id"`
	CategoryName string `json:"category_name"`
}

func (p *updateCategoryParams) mergeFieldValue(model *DBstruct.Category) {
	model.CategoryID = p.CategoryId
	model.CanteenID = p.CanteenId
	model.CategoryName = p.CategoryName
}

type delParams struct {
	CategoryId uint `json:"category_id" binding:"required"`
}

func GetCategoryList(c *gin.Context) {
	var params listParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var categorys []DBstruct.Category
	err := DBstruct.DB.Where("canteen_id = ?", params.CanteenId).Find(&categorys).Error
	if err != nil {
		c.JSON(200, api2.ERRRESPONSE("服务异常", 500))
		return
	}
	var list []vo.CategoryList
	for _, category := range categorys {
		list = append(list, vo.CategoryList{
			CategoryID:   category.CategoryID,
			CanteenId:    category.CanteenID,
			CategoryName: category.CategoryName,
		})
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"category_list": list,
		},
	})
}

func AddCategory(c *gin.Context) {
	var params addCategoryParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var category DBstruct.Category

	params.mergeFieldValue(&category)

	err := DBstruct.DB.Model(&category).Create(&category).Error
	if err != nil {
		c.JSON(200, api2.ERRRESPONSE("服务异常", 500))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE(params))
}

func UpdateCategory(c *gin.Context) {
	var params updateCategoryParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var category DBstruct.Category

	params.mergeFieldValue(&category)

	err := DBstruct.DB.Model(&category).Updates(&category).Error
	if err != nil {
		c.JSON(200, api2.ERRRESPONSE("服务异常", 500))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE(params))
}

func HasDelCategory(c *gin.Context) {
	var params delParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var product DBstruct.Product
	err := DBstruct.DB.Where("category_id = ?", params.CategoryId).First(&product).Error
	if err != nil {
		DBstruct.DB.Where("category_id = ?", params.CategoryId).Delete(&DBstruct.Category{})
		c.JSON(200, api2.SUCCESSRESPONSE(gin.H{
			"category_id": params.CategoryId,
		}))
		return
	}
	c.JSON(201, api2.SUCCESSRESPONSE("该类别中包含菜品信息，是否同步删除"))
}

func DelCategory(c *gin.Context) {
	var params delParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	DBstruct.DB.Where("category_id = ?", params.CategoryId).Delete(&DBstruct.Product{})
	DBstruct.DB.Where("category_id = ?", params.CategoryId).Delete(&DBstruct.Category{})
	c.JSON(200, api2.SUCCESSRESPONSE(gin.H{
		"category_id": params.CategoryId,
	}))
}
