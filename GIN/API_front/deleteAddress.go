package api1

import (
	"fmt"
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按照address记录的id，删除address记录
func DeleteAddress(c *gin.Context) {
	type msg struct {
		Address_id int `json:"addressID"`
		User_id    int `json:"user_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		var temp = DBstruct.Address{}
		if e := DBstruct.DB.Delete(&temp, m.Address_id).Error; e == nil {
			addressRecords := make([]DBstruct.Address, 0) //存储修改后的Address列表
			e = DBstruct.DB.Where("user_id = ?", m.User_id).Find(&addressRecords).Error
			if e == nil {
				fmt.Println(addressRecords)
				type returnmsg struct {
					Address []DBstruct.Address `json:"address"`
				}
				c.JSON(200, SUCCESSRESPONSE(returnmsg{addressRecords})) //返回修改后的地址列表
			}
		} else {
			c.JSON(201, ERRRESPONSE(e.Error(), 201)) //返回操作失败及原因
		}
	} else {
		c.JSON(400, ERRRESPONSE(e.Error(), 201)) //解析json失败
	}
}
