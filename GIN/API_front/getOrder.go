package api1

import (
	"gintest/DBstruct"
	"time"

	"github.com/gin-gonic/gin"
)

type orderJson struct {
	OrderTime time.Time `json:"order_time"`
	OrderID   uint64    `json:"order_id"` //一次下单可以选购多道菜品，每道菜品都有一条Order记录，但是有相同的OrderID，用于一起支付和发餐收餐
	// UserName     string	`json:""`
	// UserId       int    ``//下单用户的id
	// Address      string //配送的地址
	ProductId int `json:"product_id"`  //下单的产品
	Num       int `json:"product_num"` //下单的数量
	// UserPhone    string //下单用户的电话，从address数据库中获取
	Status string `json:"order_state"` //“有待支付状态，未支付的都在购物车数据库中”“未取餐”“送餐中”“已送达”
	// Canteen      int    //订单属于哪个食堂
	// DeliverID    int    //配送员编号
	// DeliverName  string //配送员姓名
	// DeliverPhone string //配送员电话
} //要加上一个product_price字段
func getJsonOrder(src DBstruct.Order) (orderJson, error) {
	product := DBstruct.Product{}
	e := DBstruct.DB.Where("id = ?", src.ProductId).First(&product).Error
	return orderJson{
		OrderTime: src.Model.CreatedAt,
		OrderID:   src.OrderID,
		ProductId: src.ProductId,
		Num:       src.Num,
		Status:    src.Status,
	}, e
}

// 读取报文体的user_id，然后按其找到该用户的所有Address记录，返回一个数组
func GetOrder(c *gin.Context) {
	type msg struct {
		UserID int `json:"user_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		orderRecords := make([]DBstruct.Order, 0)
		if e = DBstruct.DB.Where("user_id = ?", m.UserID).Find(&orderRecords).Error; e == nil {
			orderJsons := make([]orderJson, 0)
			for _, orderRecord := range orderRecords {
				if orderJson, e := getJsonOrder(orderRecord); e == nil {
					orderJsons = append(orderJsons, orderJson)
				} else {
					c.JSON(201, ERRRESPONSE(e.Error(), 201))
					return
				}
			}
			type returnmsg struct {
				OrderJsons []orderJson `json:"orders"`
			}
			c.JSON(200, SUCCESSRESPONSE(returnmsg{OrderJsons: orderJsons}))
		} else {
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
		}
	} else {
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
