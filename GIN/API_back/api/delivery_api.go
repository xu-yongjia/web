package api

import (
	api2 "gintest/API_back"
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// type getDeliveryParams struct {
// 	Page          *int   `json:"page" binding:"required" message:"page required"`
// 	NumEachPage   *int   `json:"num_each_page" binding:"required" message:"num_each_page required"`
// 	CanteenNumber int    `json:"canteen_number" binding:"required" message:"canteen_number required"`
// 	TrueName      string `json:"truename"`
// }

type addDeliveryParams struct {
	CanteenId int    `json:"canteen_id" binding:"required"`
	TrueName  string `json:"truename" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

func (p *addDeliveryParams) mergeFieldValue(model *DBstruct.Delivery) {
	model.CanteenID = p.CanteenId
	model.Truename = p.TrueName
	model.Phone = p.Phone
}

type updateDeliveryParams struct {
	DeliverId uint   `json:"deliver_id" binding:"required"`
	TrueName  string `json:"truename" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	CanteenId int    `json:"canteen_id" binding:"required"`
}

func (p *updateDeliveryParams) mergeFieldValue(model *DBstruct.Delivery) {
	model.ID = p.DeliverId
	model.CanteenID = p.CanteenId
	model.Truename = p.TrueName
	model.Phone = p.Phone
}

type delDeliveryParams struct {
	Ids []int `json:"deliver_id" binding:"required"`
}

// func GetDeliveryList(c *gin.Context) {
// 	var params getDeliveryParams
// 	if err := c.ShouldBind(&params); err != nil {
// 		c.JSON(400, api2.ERRRESPONSE("数据格式错误", 201))
// 		return
// 	}
// 	db := DBstruct.DB
// 	var delivery []DBstruct.Delivery
// 	var count int
// 	if params.TrueName != "" {
// 		db = db.Where("truename like ?", "%"+params.TrueName+"%")
// 	}

// 	err := db.Limit(*params.NumEachPage).Offset((*params.Page - 1) * (*params.NumEachPage)).Find(&delivery).Error

// 	if err != nil {
// 		c.JSON(200, api2.ERRRESPONSE("服务异常", 500))
// 		return
// 	}
// 	db.Model(&DBstruct.Delivery{}).Count(&count)
// 	var list []vo.DeliveryList
// 	for _, val := range delivery {
// 		list = append(list, vo.DeliveryList{
// 			CanteenNumber: val.CanteenID,
// 			DeliverId:     val.ID,
// 			Truename:      val.Truename,
// 			Phone:         val.Phone,
// 		})
// 	}
// 	c.JSON(200, api2.SUCCESSRESPONSE(gin.H{
// 		"count":       count,
// 		"deliverlist": list,
// 	}))
// }

func AddDelivery(c *gin.Context) {
	var params addDeliveryParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var delivery DBstruct.Delivery

	params.mergeFieldValue(&delivery)

	err := DBstruct.DB.Model(&delivery).Create(&delivery).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
}

func UpdateDelivery(c *gin.Context) {
	var params updateDeliveryParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	var delivery DBstruct.Delivery

	params.mergeFieldValue(&delivery)

	err := DBstruct.DB.Model(&delivery).Updates(&delivery).Error
	if err != nil {
		c.JSON(201, api2.ERRRESPONSE("服务异常", 201))
		return
	}
	c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
}

func DelDelivery(c *gin.Context) {
	var params delDeliveryParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(201, api2.ERRRESPONSE("数据格式错误", 201))
		return
	}
	for _, id := range params.Ids {
		DBstruct.DB.Where("id = ?", id).Delete(&DBstruct.Delivery{})
	}
	c.JSON(200, api2.SUCCESSRESPONSE_NODATA())
}
