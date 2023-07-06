package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	type msg struct {
		TargetCarts []uint `json:"carts"`
		Username    string `json:"username"`
		Userid      int    `json:"user_id"`
		Address     string `json:"address"`
		UserPhone   string `json:"phone"` //用户电话，不是配餐员电话
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		var firstCanteenId int
		newOrderID := uint64(0)
		newOrderRecords := []DBstruct.Order{}
		for _, cartID := range m.TargetCarts {
			cartRecord := DBstruct.Cart{}
			if e = DBstruct.DB.Where("id = ?", cartID).First(&cartRecord).Error; e == nil { //对于请求中的每一个cartID，查找相应的cart记录
				if firstCanteenId == 0 {
					firstCanteenId = cartRecord.CanteenID //初始化canteenid，用于判断所有记录是否对于同一个食堂
				}
				if cartRecord.CanteenID != firstCanteenId { //如果出现不同的id，报错并结束
					c.JSON(201, ERRRESPONSE("不能同时勾选不同食堂的商品！", 201))
					return
				}
				if newOrderID == 0 {
					newOrderID = uint64(cartRecord.ID) //将第一个cart记录的id作为这批order记录的order_id（主键是另一个id），避免并发处理时出错
				}
				newOrderRecord := DBstruct.Order{ //组合出新的order记录结构体
					OrderID:   newOrderID,
					UserName:  m.Username,
					UserId:    m.Userid,
					Address:   m.Address,
					ProductId: cartRecord.ProductId,
					Num:       cartRecord.Number,
					UserPhone: m.UserPhone,
					Status:    "未支付",
					CanteenID: cartRecord.CanteenID,
				}
				newOrderRecords = append(newOrderRecords, newOrderRecord) //1.不要直接修改数据库，后面可能因为食堂id不唯一而出错；2.别用冒等，作用域不同
			} else {
				c.JSON(201, ERRRESPONSE("找不到购物车记录", 201))
				return
			}
		}
		for _, new := range newOrderRecords {
			DBstruct.DB.Create(&new)
		}
		type msg struct {
			OrderID uint `json:"order_id"`
		}
		if e = DBstruct.DB.Delete(DBstruct.Cart{}, m.TargetCarts).Error; e == nil {
			c.JSON(200, SUCCESSRESPONSE(msg{OrderID: uint(newOrderID)}))
		} else {
			c.JSON(201, ERRRESPONSE("删除购物车记录时出错："+e.Error(), 201))
		}
	} else { //解析json失败
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
