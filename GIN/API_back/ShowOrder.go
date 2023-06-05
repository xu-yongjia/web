package api2

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type PaginationOrderRequest struct {
	Page        int    `json:"page"`
	NumEachPage int    `json:"num_each_page"`
	CanteenId   int    `json:"canteen_id"`
	Status      string `json:"status"`
}

type OrderDisplay struct {
	ID           uint   `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	OrderID      uint64 `json:"order_id"`
	UserName     string `json:"user_name"`
	UserId       int    `json:"user_id"`
	Address      string `json:"address"`
	ProductId    int    `json:"product_id"`
	Num          int    `json:"num"`
	UserPhone    string `json:"user_phone"`
	Status       string `json:"status"`
	Canteen      int    `json:"canteen"`
	DeliverID    int    `json:"deliver_id"`
	DeliverName  string `json:"deliver_name"`
	DeliverPhone string `json:"deliver_phone"`
}

func ShowOrder(c *gin.Context) {

	var pagination PaginationOrderRequest
	if err := c.ShouldBindJSON(&pagination); err != nil {
		c.JSON(400, gin.H{"msg": "数据格式错误", "status": 201})
		return
	}

	var orders []DBstruct.Order
	if pagination.Status != "" {
		DBstruct.DB.Where("canteen = ? AND status = ?", pagination.CanteenId, pagination.Status).Order("updated_at").Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&orders)
	} else {
		DBstruct.DB.Where("canteen = ?", pagination.CanteenId).Order("updated_at").Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&orders)
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
		orderDisplay.Address = order.Address
		orderDisplay.ProductId = order.ProductId
		orderDisplay.Num = order.Num
		orderDisplay.UserPhone = order.UserPhone
		orderDisplay.Status = order.Status
		orderDisplay.Canteen = order.Canteen
		orderDisplay.DeliverID = order.DeliverID
		orderDisplay.DeliverName = order.DeliverName
		orderDisplay.DeliverPhone = order.DeliverPhone

		orderDisplays = append(orderDisplays, orderDisplay)
	}

	var totalCount int
	if pagination.Status != "" {
		DBstruct.DB.Model(&DBstruct.Order{}).Where("canteen = ? AND status = ?", pagination.CanteenId, pagination.Status).Count(&totalCount)
	} else {
		DBstruct.DB.Model(&DBstruct.Order{}).Where("canteen = ?", pagination.CanteenId).Count(&totalCount)
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

