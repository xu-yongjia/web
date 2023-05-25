package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按照用户ID和商品ID，从购物车中删除商品
func DeleteShoppingCart(c *gin.Context) {
	type msg struct {
		User_id    int `json:"user_id"`
		Product_id int `json:"product_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		if e := DBstruct.DB.Where("user_id=? AND product_id=?", m.User_id, m.Product_id).Delete(&DBstruct.Cart{}).Error; e == nil { //成功删除
			c.JSON(200, SUCCESSRESPONSE_NODATA())
		} else { //删除出错（如果不存在记录，不会报错）
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
		}
	} else { //json解析出错
		c.JSON(201, ERRRESPONSE(e.Error(), 201))
	}
}
