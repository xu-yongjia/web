package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// User 商品序列化器
type Count_json struct {
	FavoriteCount int `json:"favorite_total"`
	NotPayCount   int `json:"not_pay_total"` //这里有一个“电话”信息，但是系统利用的电话信息主要是从Address记录中获取的，与这里的无关，只是为了保证扩展性而添加的冗余字段
	PayCount      int `json:"pay_total"`
	WaitJudge     int `json:"wait_judge"`
}

// 验证发来的用户名和密码（可能报告账号不存在或者密码错误），并生成一个token，随完整用户数据一起返回
func ShowCount(c *gin.Context) {
	type msg struct {
		UserID int `json:"user_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		var favoriteCount, notPayCount, payCount, waitJudge int
		DBstruct.DB.Model(&DBstruct.Favorite{}).Where("user_id = ?", m.UserID).Count(&favoriteCount)
		DBstruct.DB.Model(&DBstruct.Order{}).Where("status = '未支付' AND user_id = ?", m.UserID).Count(&notPayCount)
		DBstruct.DB.Model(&DBstruct.Order{}).Where("status <> '未支付' AND user_id = ?", m.UserID).Count(&payCount)
		DBstruct.DB.Model(&DBstruct.Order{}).Where("status = '已送达' AND user_id = ?", m.UserID).Count(&waitJudge)
		c.JSON(200, SUCCESSRESPONSE(Count_json{FavoriteCount: favoriteCount, NotPayCount: notPayCount, PayCount: payCount, WaitJudge: waitJudge}))
	} else { //JSON解析失败
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
