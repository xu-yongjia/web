package api1

import (
	"fmt"
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 读取报文体的user_id，然后按其找到该用户的所有Address记录，返回一个数组
func GetUserAddress(c *gin.Context) {
	type msg struct {
		UserID int `json:"user_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		addressRecords := make([]DBstruct.Address, 0)
		if e = DBstruct.DB.Where("user_id = ?", m.UserID).Find(&addressRecords).Error; e == nil {
			fmt.Println(addressRecords)
			type returnmsg struct { //包装data
				Address []DBstruct.Address `json:"address"`
			}
			c.JSON(200, SUCCESSRESPONSE(returnmsg{addressRecords}))
		}
	} else {
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
