package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

// 把外键原本地返回，没有做二次查找，后续可能要加上（canttenID和categoryID）
func GetDetails(c *gin.Context) {
	type msg struct {
		ProductID int `json:"productID"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		var Productmsg []DBstruct.Product //因为返回时需要以数组的格式，所以用slice存储查找结果，否则用一个Product结构体即可
		if e = DBstruct.DB.Where("id = ?", m.ProductID).First(&Productmsg).Error; e == nil {
			type returnData struct { //包装返回data
				ProductSlice []DBstruct.Product `json:"Product"`
			}
			c.JSON(200, SUCCESSRESPONSE(returnData{Productmsg})) //返回商品详细信息
		} else {
			c.JSON(201, ERRRESPONSE("商品记录不存在", 201)) //报错(使用First查找记录时，如果记录不存在，会报错)
		}
	} else { //JSON解析出错
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
