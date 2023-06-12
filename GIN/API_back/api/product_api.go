package api

import (
	api2 "gintest/API_back"
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type productJson struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Category_id    uint   `json:"category_id"`
	Category_name  string `json:"category_name"`
	Info           string `json:"Info"`
	Img_path       string `json:"img_path"`
	Price          string `json:"price"`
	Discount_price string `json:"discount_price"`
	Title          string `json:"titile"`
	Score          string `json:"score"`
	Sales          int    `json:"sales"`
}

func Buildproducts(a []DBstruct.Product) (b []productJson, e error) {
	for _, product := range a {
		category := DBstruct.Category{CategoryID: uint(product.CategoryID)}
		e = DBstruct.DB.First(&category).Error
		if e != nil {
			return b, e
		}
		b = append(b, productJson{
			ID:             product.ID,
			Name:           product.Name,
			Category_id:    product.CategoryID,
			Category_name:  category.CategoryName,
			Info:           product.Info,
			Img_path:       product.ImgPath,
			Price:          product.Price,
			Discount_price: product.DiscountPrice,
			Title:          product.Title,
			Score:          product.Score,
			Sales:          product.Sales,
		})
	}
	return b, e
}

type productListParams struct {
	Page       *int `json:"page" binding:"required"`
	Limit      *int `json:"limit" binding:"required"`
	CanteenId  int  `json:"canteen_id"`
	CategoryId int  `json:"category_id"`
}

type addProductParams struct {
	Name          string `json:"name" binding:"required"`
	CanteenId     int    `json:"canteen_id" binding:"required"`
	CategoryId    uint   `json:"category_id" binding:"required"`
	Info          string `json:"info" binding:"required"`
	Price         string `json:"price" binding:"required"`
	DiscountPrice string `json:"discount_price" binding:"required"`
	Title         string `json:"title" binding:"required"`
}

func (p *addProductParams) mergeFieldValue(model *DBstruct.Product) {
	model.Name = p.Name
	model.CanteenID = p.CanteenId
	model.CategoryID = p.CategoryId
	model.Info = p.Info
	model.Price = p.Price
	model.DiscountPrice = p.DiscountPrice
	model.Title = p.Title
}

type updateProductParams struct {
	Id            uint   `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	CanteenId     int    `json:"canteen_id" binding:"required"`
	CategoryId    int    `json:"category_id" binding:"required"`
	Info          string `json:"info" binding:"required"`
	Price         string `json:"price" binding:"required"`
	DiscountPrice string `json:"discount_price" binding:"required"`
	Title         string `json:"title" binding:"required"`
}

func (p *updateProductParams) mergeFieldValue(model *DBstruct.Product) {
	model.ID = p.Id
	model.Name = p.Name
	model.CanteenID = p.CanteenId
	model.CategoryID = uint(p.CategoryId)
	model.Info = p.Info
	model.Price = p.Price
	model.DiscountPrice = p.DiscountPrice
	model.Title = p.Title
}

type delProductParams struct {
	ProductId int `json:"product_id" binding:"required"`
}

// "product_id":123,
// "canteen_id":1
// "big_url":"http://...",
// "small_urls":["http://...","http://...","http://..."]
type savePhotoParams struct {
	ProductId int      `json:"product_id" binding:"required"`
	CanteenId int      `json:"canteen_id" binding:"required"`
	BigUrl    string   `json:"big_url" binding:"required"`
	SmallUrls []string `json:"small_urls" binding:"required"`
}

func AddProduct(c *gin.Context) {
	var params addProductParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var product DBstruct.Product

	params.mergeFieldValue(&product)

	err := DBstruct.DB.Model(&product).Create(&product).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
}

func UpdateProduct(c *gin.Context) {
	var params updateProductParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var product DBstruct.Product

	params.mergeFieldValue(&product)

	err := DBstruct.DB.Model(&product).Updates(&product).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
}

func GetProductList(c *gin.Context) {
	var params productListParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}

	db := DBstruct.DB
	if params.CategoryId != 0 {
		db = db.Where("category_id = ?", params.CategoryId)
	}
	if params.CanteenId != 0 {
		db = db.Where("canteen_id = ?", params.CanteenId)
	}

	var count int
	var products []DBstruct.Product
	var products_json []productJson
	err := db.Limit(*params.Limit).Offset((*params.Page - 1) * (*params.Limit)).Find(&products).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	if products_json, err = Buildproducts(products); err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	db.Model(&DBstruct.Product{}).Count(&count)

	c.JSON(200, api2.SUCCESSRESPONSE(gin.H{
		"count":       count,
		"productlist": products_json,
	}))
}

func DelProduct(c *gin.Context) {
	var params delProductParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var order DBstruct.Order
	err := DBstruct.DB.Where("product_id = ? and `status` <> ?", params.ProductId, "已送达").First(&order).Error
	if err != nil {
		DBstruct.DB.Where("id = ?", params.ProductId).Delete(&DBstruct.Product{})
		//DBstruct.DB.Where("product_id = ?", params.ProductId).Delete(&DBstruct.Product{})
		c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
		return
	}
	c.JSON(201, api2.ERRRESPONSE("有未完成的该商品订单，删除失败", 201))
}

func SavePhoto(c *gin.Context) {
	var params savePhotoParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	err := DBstruct.DB.Model(&DBstruct.Product{}).Where("id = ?", params.ProductId).Update("img_path", params.BigUrl).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	for _, url := range params.SmallUrls {
		DBstruct.DB.Model(&DBstruct.ProductImg{}).Create(&DBstruct.ProductImg{ProductID: params.ProductId, ImgPath: url})
	}
	c.JSON(200, api2.SUCCESSRESPONSE(gin.H{}))
}
