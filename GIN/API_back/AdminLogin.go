package api2

import (
	"fmt"
	"gintest/DBstruct"
	"gintest/pkg/util"

	"github.com/gin-gonic/gin"
)

// AdminLogin 验证发来的管理员ID和密码（可能报告账号不存在或者密码错误），并生成一个token，随完整管理员数据一起返回
func AdminLogin(c *gin.Context) {
	fmt.Print("here!!!")
	type msg struct {
		CanteenID int    `json:"canteen_id"`
		Password  string `json:"password"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		fmt.Println("here")
		var adminRecord DBstruct.Canteen
		if e = DBstruct.DB.Where("canteen_id = ?", m.CanteenID).First(&adminRecord).Error; e == nil { //按照CanteenID查找账号
			if adminRecord.Password == m.Password { //账号密码和请求中的密码相同
				token, _ := util.GenerateToken(fmt.Sprint(adminRecord.CanteenID), adminRecord.Password, 0) //用CanteenID和Password生成一个token，主要是为了确保同一时刻产生的token不会重复
				type token_temp struct {
					Token string `json:"token"`
				}
				c.JSON(200, SUCCESSRESPONSE(token_temp{Token: token})) //返回登录成功消息
			} else {
				c.JSON(201, ERRRESPONSE("密码错误", 201)) //账号密码和请求中的密码不同
			}
		} else {
			c.JSON(201, ERRRESPONSE("账号不存在"+e.Error(), 201)) //具有请求中的CanteenID的账号不存在
		}
	} else { //JSON解析失败
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
