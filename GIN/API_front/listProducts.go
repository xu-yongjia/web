package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// DataList 带有总数的Data结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// ListProductsService 视频列表服务
type ListProductsService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
	// CanteenID  uint `form:"canteen_id" json:"canteen_id"`
	CanteenID uint `form:"placeID" json:"placeID"`
	// CategoryID uint `form:"category_id" json:"category_id"`
	CategoryID uint `form:"categoryID" json:"categoryID"`
}

// List 视频列表
func (service *ListProductsService) List() EResponse {
	products := []DBstruct.Product{}

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	if service.CanteenID == 0 {
		if err := DBstruct.DB.Model(DBstruct.Product{}).Count(&total).Error; err != nil {
			code = e.ERROR_DATABASE
			return EResponse{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			code = e.ERROR_DATABASE
			return EResponse{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	if service.CategoryID == 0 { //返回全部
		if err := DBstruct.DB.Model(DBstruct.Product{}).Where("canteen_id=?", service.CanteenID).Count(&total).Error; err != nil {
			code = e.ERROR_DATABASE
			return EResponse{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Where("canteen_id=?", service.CanteenID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			code = e.ERROR_DATABASE
			return EResponse{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := DBstruct.DB.Model(DBstruct.Product{}).Where("category_id=? AND canteen_id=?", service.CategoryID, service.CanteenID).Count(&total).Error; err != nil {
			code = e.ERROR_DATABASE
			return EResponse{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := DBstruct.DB.Where("category_id=? AND canteen_id=?", service.CategoryID, service.CanteenID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			code = e.ERROR_DATABASE
			return EResponse{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return BuildListResponse(BuildProducts(products), uint(total))
}

// ListProducts 商品列表接口
func ListProducts(c *gin.Context) {
	var service ListProductsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ERRRESPONSE(err.Error(), 201))
	}

}

// Product 商品序列化器
type Product struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	CanteenID     int     `json:"canteen_id"`
	CategoryID    int     `json:"category_id"`
	Title         string  `json:"title"`
	Info          string  `json:"info"`
	ImgPath       string  `json:"img_path"`
	Price         float32 `json:"price"`
	DiscountPrice float32 `json:"discount_price"`
	View          uint64  `json:"view"`
	CreatedAt     int64   `json:"created_at"`
}

// BuildProduct 序列化商品
func BuildProduct(item DBstruct.Product) Product {
	return Product{
		ID:            item.ID,
		Name:          item.Name,
		CanteenID:     item.CanteenID,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		CreatedAt:     item.CreatedAt.Unix(),
	}
}

// BuildProducts 序列化商品列表
func BuildProducts(items []DBstruct.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
