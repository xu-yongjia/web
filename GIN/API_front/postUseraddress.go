package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 从请求体中读取一个Address记录（不读取id，否则可能主键重复），然后确认在数据库中没有重复的记录（除了id外的其他属性全部相同），将其插入，再将具有相同user_id的所有Address记录返回
func PostUseraddress(c *gin.Context) {
	m := DBstruct.Address{}
	if e := c.ShouldBindJSON(&m); e == nil {
		if e = DBstruct.DB.Where("user_id = ? AND user_name = ? AND phone = ? AND address = ?", m.User_id, m.UserName, m.Phone, m.Address).First(&DBstruct.Address{}).Error; e == nil { //确认没有重复记录
			c.JSON(201, ERRRESPONSE("已有内容相同的记录", 201)) //如果有重复记录，则报错并结束
			return
		}
		if e = DBstruct.DB.Omit("id").Create(&m).Error; e == nil { //插入成功
			addressRecords := make([]DBstruct.Address, 0) //存储返回的Address数组
			e = DBstruct.DB.Where("user_id = ?", m.User_id).Find(&addressRecords).Error
			if e == nil {
				type returnmsg struct {
					Address []DBstruct.Address `json:"address"`
				}
				c.JSON(200, SUCCESSRESPONSE(returnmsg{addressRecords})) //返回修改后的地址列表
			}
		} else { //插入失败
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
		}
	} else {
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
