package DBstruct

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserId    int //下单用户的id
	ProductId int //下单的产品
	Number    int //下单的数量
	CanteenID int //订单属于哪个食堂
}
