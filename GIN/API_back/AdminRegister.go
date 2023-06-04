package api2


import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按照请求中的username和password产生用户记录，userid在插入记录时由数据库分配，不能修改（如果和已有用户重名则不能成功注册，因为username被声明为unique字段）
func AdminRegister(c *gin.Context) {
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
			c.JSON(200, gin.H{"status": 200, "message": "success", "data": nil})
		} else { //报错
			c.JSON(201, gin.H{"status": 201, "message": e.Error(), "data": nil})
		}
	} else { //解析json失败
		c.JSON(400, gin.H{"status": 400, "message": e.Error(), "data": nil})
	}
}
