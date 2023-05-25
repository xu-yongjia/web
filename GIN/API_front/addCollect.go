package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按照用户ID和商品ID，添加新的“收藏”记录
func AddCollect(c *gin.Context) {
	type msg struct {
		User_id    int `json:"user_id"`
		Product_id int `json:"product_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		newCollectRecord := DBstruct.Favorite{
			UserID:   m.User_id,
			ProducID: m.Product_id,
		}
		if e = DBstruct.DB.Create(&newCollectRecord).Error; e == nil { //成功添加
			c.JSON(200, SUCCESSRESPONSE_NODATA())
		} else { //报错
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
		}
	} else { //解析json失败
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
