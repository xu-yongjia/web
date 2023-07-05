package api1

import (
	"fmt"
	"gintest/DBstruct"

	"gintest/pkg/util"

	"github.com/gin-gonic/gin"
)

// User 商品序列化器
type User_json struct {
	ID       uint   `json:"user_id"`
	Phone    string `json:"phone"` //这里有一个“电话”信息，但是系统利用的电话信息主要是从Address记录中获取的，与这里的无关，只是为了保证扩展性而添加的冗余字段
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}

func BuildUser(u DBstruct.User) User_json {
	return User_json{
		ID:       u.ID,
		Phone:    u.Phone,
		UserName: u.UserName,
	}
}

// 验证发来的用户名和密码（可能报告账号不存在或者密码错误），并生成一个token，随完整用户数据一起返回
func UserLogin(c *gin.Context) {
	type msg struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		fmt.Println("here")
		var userRecord DBstruct.User
		if e = DBstruct.DB.Where("user_name = ?", m.Username).First(&userRecord).Error; e == nil { //按照username查找账号
			if userRecord.Password == m.Password { //账号密码和请求中的密码相同
				type returnmsg struct { //包装返回的数据：一个token字符串和一个user结构体
					User  User_json `json:"user"`
					Token string    `json:"token"`
				}
				token, _ := util.GenerateToken(userRecord.UserName, userRecord.Password, 0) //用username和Password生成一个token，主要是为了确保同一时刻产生的token不会重复
				c.JSON(200, SUCCESSRESPONSE(returnmsg{BuildUser(userRecord), token}))       //返回登录成功消息
			} else {
				c.JSON(201, ERRRESPONSE("密码错误", 201)) //账号密码和请求中的密码不同
			}
		} else {
			c.JSON(201, ERRRESPONSE("账号不存在"+e.Error(), 201)) //具有请求中的Username的站好不存在
		}
	} else { //JSON解析失败
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
