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
	var category DBstruct.Category
	err := DBstruct.DB.Where("category_name = ?", p.CategoryName).Find(&category).Error
	if err == nil {
		model.CategoryID = category.CategoryID
	}
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
	CanteenId  uint `json:"canteen_id" binding:"required"`
}

func GetCategoryList(c *gin.Context) {
	var params listParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var categorys []DBstruct.Category
	err := DBstruct.DB.Where("canteen_id = ?", params.CanteenId).Find(&categorys).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
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

	c.JSON(200, api2.SUCCESSRESPONSE(gin.H{"category_list": list}))
}

func AddCategory(c *gin.Context) {
	var params addCategoryParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var category DBstruct.Category

	params.mergeFieldValue(&category)

	err := DBstruct.DB.Model(&category).Create(&category).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
}

func UpdateCategory(c *gin.Context) {
	var params updateCategoryParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var category DBstruct.Category

	params.mergeFieldValue(&category)

	err := DBstruct.DB.Model(&category).Updates(&category).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
}

func HasDelCategory(c *gin.Context) {
	var params delParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var product DBstruct.Product
	err := DBstruct.DB.Where("category_id = ? AND canteen_id = ?", params.CategoryId, params.CanteenId).First(&product).Error //查找该分类下的产品
	if err != nil {                                                                                                           //如果不存在该分类下的产品，则删除
		DBstruct.DB.Where("category_id = ? AND canteen_id = ?", params.CategoryId, params.CanteenId).Delete(&DBstruct.Category{})
		c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
		return
	}
	c.JSON(201, api2.ERRRESPONSE("该类别中包含菜品信息，是否同步删除?", 201))
}

func DelCategory(c *gin.Context) {
	var params delParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	DBstruct.DB.Where("category_id = ? AND canteen_id = ?", params.CategoryId, params.CanteenId).Delete(&DBstruct.Product{})
	DBstruct.DB.Where("category_id = ? AND canteen_id = ?", params.CategoryId, params.CanteenId).Delete(&DBstruct.Category{})
	c.JSON(201, api2.SUCCESSRESPONSE_NODATA())
}
