package api1

import (
	"fmt"
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 按照用户ID和商品ID，添加新的“收藏”记录
func GetCollect(c *gin.Context) {
	type msg struct {
		User_id int `json:"user_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		var Favorite []DBstruct.Favorite
		e := DBstruct.DB.Where("user_id = ?", m.User_id).Order("updated_at").Find(&Favorite).Error
		if e != nil {
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
			return
		}
		var product_json []Product_json
		for _, favorite := range Favorite {
			productId := favorite.ProducID
			var product DBstruct.Product
			if err := DBstruct.DB.First(&product, productId).Error; err != nil {
				c.JSON(201, ERRRESPONSE(e.Error(), 201))
				return
			}
			product_json = append(product_json, BuildProduct(product))
			fmt.Print(product_json)
		}
		c.JSON(200, gin.H{
			"msg":    "OK",
			"status": 200,
			"data": gin.H{
				"products": product_json,
			},
		})
	} else { //解析json失败
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
