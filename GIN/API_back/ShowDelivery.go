package api2

import (
	"gintest/DBstruct"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeliveryPaginationRequest struct {
	Page         int    `json:"page"`
	NumEachPage  int    `json:"num_each_page"`
	CanteenID    int    `json:"canteen_id"`
	TruenameOrID string `json:"truename_or_id"`
}

type DeliveryDisplay struct {
	DeliverID uint   `json:"deliver_id"`
	Truename  string `json:"truename"`
	Phone     string `json:"phone"`
	CanteenID int    `json:"canteen_id"`
}

func ShowDelivery(c *gin.Context) {
	var pagination DeliveryPaginationRequest
	if err := c.ShouldBindJSON(&pagination); err != nil {
		c.JSON(201, ERRRESPONSE("json参数格式有误", 201))
		return
	}

	var deliveries []DBstruct.Delivery
	var totalDeliveries []DBstruct.Delivery
	var total int

	DBstruct.DB.Where("canteen_id = ?", pagination.CanteenID).Find(&totalDeliveries)
	total = len(totalDeliveries)

	if pagination.TruenameOrID != "" {
		if id, err := strconv.Atoi(pagination.TruenameOrID); err == nil {
			DBstruct.DB.Where("canteen_id = ? AND (truename like ? OR id = ?)", pagination.CanteenID, "%"+pagination.TruenameOrID+"%", id).Find(&deliveries)
		} else {
			DBstruct.DB.Where("canteen_id = ? AND truename like ?", pagination.CanteenID, "%"+pagination.TruenameOrID+"%").Find(&deliveries)
		}
	} else {
		DBstruct.DB.Where("canteen_id = ?", pagination.CanteenID).Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&deliveries)
	}

	var deliverDisplays []DeliveryDisplay
	for _, delivery := range deliveries {
		var deliverDisplay DeliveryDisplay
		deliverDisplay.DeliverID = delivery.ID
		deliverDisplay.Truename = delivery.Truename
		deliverDisplay.Phone = delivery.Phone
		deliverDisplay.CanteenID = delivery.CanteenID

		deliverDisplays = append(deliverDisplays, deliverDisplay)
	}

	c.JSON(200, SUCCESSRESPONSE(gin.H{
		"count":       total,
		"deliverlist": deliverDisplays,
	}))
}
