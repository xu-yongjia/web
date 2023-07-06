package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// // User 商品序列化器
// type Count_json struct {
// 	FavoriteCount int `json:"favorite_total"`
// 	NotPayCount   int `json:"not_pay_total"` //这里有一个“电话”信息，但是系统利用的电话信息主要是从Address记录中获取的，与这里的无关，只是为了保证扩展性而添加的冗余字段
// 	PayCount      int `json:"pay_total"`
// 	WaitJudge     int `json:"wait_judge"`
// }

// 验证发来的用户名和密码（可能报告账号不存在或者密码错误），并生成一个token，随完整用户数据一起返回
func RecieveOrder(c *gin.Context) {
	type msg struct {
		OrderID int `json:"order_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		DBstruct.DB.Model(DBstruct.Order{}).Where("order_id = ?", m.OrderID).Update("status", "已送达")
		c.JSON(200, Response{
			Status: 200,
			Msg:    "收餐成功",
			Data:   "",
		})
	} else { //JSON解析失败
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
