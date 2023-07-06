package api1

import (
	"fmt"
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type Json_address struct {
	ID       int    `json:"ID"`
	User_id  int    `json:"user_id"`
	UserName string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Index    int    `json:"index"`
}

// 读取报文体的user_id，然后按其找到该用户的所有Address记录，返回一个数组
func GetUserAddress(c *gin.Context) {
	type msg struct {
		UserID int `json:"user_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		var addressRecords []DBstruct.Address
		var json_addresss []Json_address
		if e = DBstruct.DB.Where("user_id = ?", m.UserID).Find(&addressRecords).Error; e == nil {
			var index = 1
			for _, address := range addressRecords {
				json_addresss = append(json_addresss, Json_address{
					ID:       int(address.ID),
					UserName: address.UserName,
					User_id:  address.User_id,
					Phone:    address.Phone,
					Address:  address.Address,
					Index:    index})
				index++
			}

			fmt.Println(addressRecords)
			type returnmsg struct { //包装data
				Address []Json_address `json:"address"`
			}
			c.JSON(200, SUCCESSRESPONSE(returnmsg{json_addresss}))
		}
	} else {
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
