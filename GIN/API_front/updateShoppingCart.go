package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按主键找到json中包含的Address数据响应的记录，然后将记录替换为json中的Address，并返回新的完整列表
// 这里的实现效果是：如果发来的Address中有空字段，那么字段也会被替换为空字段
func UpdateShoppingCart(c *gin.Context) {
	type msg struct {
		User_id    int `json:"user_id"`
		Num        int `json:"num"`
		Product_id int `json:"product_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		cartRecord := DBstruct.Cart{}
		if e = DBstruct.DB.Where("user_id=? AND product_id=?", m.User_id, m.Product_id).First(&cartRecord).Error; e == nil { //根据userID和productID查找原来的Address记录
			cartRecord.Number = m.Num
			if e = DBstruct.DB.Save(&cartRecord).Select("number").Error; e == nil { //修改数据库中Address记录的Number字段
				c.JSON(200, SUCCESSRESPONSE_NODATA())
			} else {
				c.JSON(201, ERRRESPONSE(e.Error(), 201)) //修改失败
			}
		} else {
			c.JSON(201, ERRRESPONSE(e.Error(), 201)) //要修改的Address记录不存在
		}
	} else { //JSON解析失败
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
