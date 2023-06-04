package api1

import (
	"gintest/DBstruct"
	"time"

	"github.com/gin-gonic/gin"
)

type cart_json struct {
	ID            uint      `json:"cart_id"`
	UserId        int       `json:"user_id"`
	ProductId     int       `json:"product_id"` //下单的产品
	Number        int       `json:"number"`     //下单的数量
	CanteenID     int       `json:"canteen_id"` //订单属于哪个食堂
	Price         string    `json:"price"`
	DiscountPrice string    `json:"discount_price"` //折后价
	Info          string    `json:"product_intro"`  //商品简介
	ImgPath       string    `json:"img_path"`       //外面的大图
	CreateAt      time.Time `json:"create_at"`      //添加购物车的时间
	Title         string    `json:"title"`          //标题
	Score         string    `json:"score"`          //评分
}

func getCartsJson(src DBstruct.Cart) (cart_json, error) {
	product := DBstruct.Product{}
	e := DBstruct.DB.Where("id = ?", src.ProductId).First(&product).Error
	return cart_json{
		ID:            src.ID,
		UserId:        src.UserId,
		ProductId:     src.ProductId,
		Number:        src.Number,
		CanteenID:     src.CanteenID,
		Price:         product.Price,
		DiscountPrice: product.Price,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		CreateAt:      product.CreatedAt,
		Title:         product.Title,
		Score:         product.Score,
	}, e
}

// 读取报文体的user_id，然后按其找到该用户的所有Address记录，返回一个数组
func GetCart(c *gin.Context) {
	type msg struct {
		UserID int `json:"user_id"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		CartRecords := make([]DBstruct.Cart, 0)
		if e = DBstruct.DB.Where("user_id = ?", m.UserID).Find(&CartRecords).Error; e == nil {
			CartJsons := make([]cart_json, 0)
			for _, cartRecord := range CartRecords {
				if cartJson, e := getCartsJson(cartRecord); e == nil {
					CartJsons = append(CartJsons, cartJson)
				} else {
					c.JSON(201, ERRRESPONSE(e.Error(), 201))
					return
				}
			}
			type returnmsg struct {
				OrderJsons []cart_json `json:"carts"`
			}
			c.JSON(200, SUCCESSRESPONSE(returnmsg{OrderJsons: CartJsons}))
		} else {
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
		}
	} else {
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
