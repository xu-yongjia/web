package api1

import (
	"gintest/DBstruct"

	"strconv"

	"github.com/gin-gonic/gin"
)

// 按照User_id Product_id添加购物车记录，其中商品所属食堂通过查找数据表“product”获取，商品数量默认为1
func AddShoppingCart(c *gin.Context) {
	type msg struct {
		User_id    int    `json:"user_id"`
		Product_id string `json:"product_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		if e = DBstruct.DB.Where("user_id=? AND product_id=?", m.User_id, m.Product_id).First(&DBstruct.Cart{}).Error; e == nil { //禁止重复添加购物车
			c.JSON(201, ERRRESPONSE("已存在于购物车中", 201))
			return
		}
		productRecord := DBstruct.Product{}
		if e = DBstruct.DB.Where("id = ? ", m.Product_id).First(&productRecord).Error; e == nil { //查找商品记录，获取食堂ID等信息
			productid, _ := strconv.Atoi(m.Product_id)
			newCartRecord := DBstruct.Cart{ //新的购物车记录
				UserId:    m.User_id,
				ProductId: productid,
				Number:    1,
				CanteenID: productRecord.CanteenID,
			}
			if e = DBstruct.DB.Create(&newCartRecord).Error; e == nil { //插入成功
				shoppingCartData := make([]DBstruct.Cart, 0)
				DBstruct.DB.Where("user_id = ?", m.User_id).Order("id desc").Find(&shoppingCartData) //按主键倒序查找，将最大的id排在最前面
				type returnmsg struct {                                                              //包装返回信息
					CartData []DBstruct.Cart `json:"shoppingCartData"`
				}
				c.JSON(200, SUCCESSRESPONSE(returnmsg{shoppingCartData})) //返回修改后的地址列表
			} else { //插入失败
				c.JSON(201, ERRRESPONSE(e.Error(), 201))
			}
		} else { //解析json失败
			c.JSON(400, ERRRESPONSE(e.Error(), 201))
		}
	} else {
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}
