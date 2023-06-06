package api2

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type OrderProduct struct {
	Id            uint   `json:"id"`                    // Order ID corresponds to Order struct
	ProductId     uint   `json:"product_id"`            // Product ID corresponds to Order struct
	ProductName   string `json:"product_name"`          // Product name corresponds to Product struct
	ImgPath       string `json:"img_path"`              // ImgPath corresponds to Product struct
	Num           int    `json:"num"`                   // Num corresponds to Order struct
	Price         string `json:"product_selling_price"` // Price corresponds to Product struct
	DiscountPrice string `json:"product_price"`         // DiscountPrice corresponds to Product struct
}

type PaginationOrderRequest struct {
	Page        int    `json:"page"`
	NumEachPage int    `json:"num_each_page"`
	CanteenId   int    `json:"canteen_id"`
	Status      string `json:"status"`
}

type OrderDisplay struct {
	ID               uint           `json:"id"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
	OrderID          uint64         `json:"order_id"`
	UserName         string         `json:"user_name"`
	UserId           int            `json:"user_id"`
	Address          string         `json:"address"`
	UserPhone        string         `json:"user_phone"`
	Status           string         `json:"status"`
	CanteenID        int            `json:"canteen_id"`
	DeliverID        int            `json:"deliver_id"`
	DeliverName      string         `json:"deliver_name"`
	DeliverPhone     string         `json:"deliver_phone"`
	OrderProductList []OrderProduct `json:"order_productlist"`
}

func ShowOrder(c *gin.Context) {
	var pagination PaginationOrderRequest
	if err := c.ShouldBindJSON(&pagination); err != nil {
		c.JSON(400, gin.H{"msg": "数据格式错误", "status": 201})
		return
	}

	var orders []DBstruct.Order
	if pagination.Status != "" {
		DBstruct.DB.Where("canteen_id = ? AND status = ?", pagination.CanteenId, pagination.Status).Order("updated_at").Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&orders)
	} else {
		DBstruct.DB.Where("canteen_id = ?", pagination.CanteenId).Order("updated_at").Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&orders)
	}

	var orderDisplays []OrderDisplay
	for _, order := range orders {
		var orderDisplay OrderDisplay
		orderDisplay.ID = order.ID
		orderDisplay.CreatedAt = order.CreatedAt.String()
		orderDisplay.UpdatedAt = order.UpdatedAt.String()
		orderDisplay.OrderID = order.OrderID
		orderDisplay.UserName = order.UserName
		orderDisplay.UserId = order.UserId
		orderDisplay.Status = order.Status
		orderDisplay.CanteenID = order.CanteenID
		orderDisplay.DeliverID = order.DeliverID
		orderDisplay.DeliverName = order.DeliverName
		orderDisplay.DeliverPhone = order.DeliverPhone

		// Find address and phone of user from Address table
		var address DBstruct.Address
		DBstruct.DB.Where("user_id = ?", order.UserId).First(&address)
		orderDisplay.Address = address.Address
		orderDisplay.UserPhone = address.Phone

		// Find order products associated with each order
		var product DBstruct.Product
		DBstruct.DB.Where("id = ?", order.ProductId).First(&product)

		var orderProduct OrderProduct
		orderProduct.Id = order.ID
		orderProduct.ProductId = product.ID
		orderProduct.ProductName = product.Name
		orderProduct.ImgPath = product.ImgPath
		orderProduct.Num = order.Num
		orderProduct.Price = product.Price
		orderProduct.DiscountPrice = product.DiscountPrice

		// Add order product to order's product list
		orderDisplay.OrderProductList = append(orderDisplay.OrderProductList, orderProduct)

		orderDisplays = append(orderDisplays, orderDisplay)
	}

	var totalCount int
	if pagination.Status != "" {
		DBstruct.DB.Model(&DBstruct.Order{}).Where("canteen_id = ? AND status = ?", pagination.CanteenId, pagination.Status).Count(&totalCount)
	} else {
		DBstruct.DB.Model(&DBstruct.Order{}).Where("canteen_id = ?", pagination.CanteenId).Count(&totalCount)
	}

	c.JSON(200, gin.H{
		"msg":    "OK",
		"status": 200,
		"data": gin.H{
			"count":     totalCount,
			"orderlist": orderDisplays,
		},
	})
}
