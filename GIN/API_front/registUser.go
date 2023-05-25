package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按照请求中的username和password产生用户记录，userid在插入记录时由数据库分配，不能修改（如果和已有用户重名则不能成功注册，因为username被声明为unique字段）
func RegistUser(c *gin.Context) {
	type msg struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		newUserRecord := DBstruct.User{
			UserName: m.UserName,
			Password: m.Password,
		}
		if e = DBstruct.DB.Create(&newUserRecord).Error; e == nil {
			c.JSON(200, SUCCESSRESPONSE_NODATA())
		} else { //报错
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
		}
	} else { //解析json失败
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
