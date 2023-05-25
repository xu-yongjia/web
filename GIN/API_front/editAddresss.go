package api1

import (
	"fmt"
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按主键找到json中包含的Address数据响应的记录，然后将记录替换为json中的Address，并返回新的完整列表
// 这里的实现效果是：如果发来的Address中有空字段，那么响应的字段也会被替换为空字段
func EditAddress(c *gin.Context) {
	m := DBstruct.Address{}
	if e := c.ShouldBindJSON(&m); e == nil {
		fmt.Println("收到参数：", m)
		if e = DBstruct.DB.Model(m).Updates(m).Error; e == nil { //修改成功
			addressRecords := make([]DBstruct.Address, 0) //存储修改后的Address列表
			e = DBstruct.DB.Where("user_id = ?", m.User_id).Find(&addressRecords).Error
			if e == nil {
				fmt.Println(addressRecords)
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
