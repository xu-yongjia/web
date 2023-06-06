package api2

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type AssignDeliveryRequest struct {
	OrderIDList []int `json:"order_id_list"`
	DeliverID   int   `json:"deliver_id"`
}

func AssignDelivery(c *gin.Context) {
	var assignDeliveryRequest AssignDeliveryRequest
	if err := c.ShouldBindJSON(&assignDeliveryRequest); err != nil {
		c.JSON(400, gin.H{"msg": "数据格式错误", "status": 201})
		return
	}

	var delivery DBstruct.Delivery
	result := DBstruct.DB.First(&delivery, assignDeliveryRequest.DeliverID)

	if result.Error != nil {
		c.JSON(201, gin.H{"msg": "无法找到指定的配送员", "status": 201})
		return
	}

	for _, orderID := range assignDeliveryRequest.OrderIDList {
		var order DBstruct.Order
		result := DBstruct.DB.First(&order, orderID)
		if result.Error != nil {
			c.JSON(201, gin.H{"msg": "无法找到指定的订单", "status": 201})
			return
		}

		if order.Status != "已支付" {
			c.JSON(201, gin.H{"msg": "订单状态必须为'已支付'", "status": 201})
			return
		}

		order.Status = "送餐中"
		order.DeliverID = assignDeliveryRequest.DeliverID
		order.DeliverName = delivery.Truename
		order.DeliverPhone = delivery.Phone
		DBstruct.DB.Save(&order)
	}

	c.JSON(200, gin.H{
		"msg":    "OK",
		"status": 200,
		"data":   gin.H{},
	})
}
