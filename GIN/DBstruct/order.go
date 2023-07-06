package DBstruct

import "github.com/jinzhu/gorm"

// Address 收货地址模型
type Order struct {
	gorm.Model
	OrderID      uint64 //一次下单可以选购多道菜品，每道菜品都有一条Order记录，但是有相同的OrderID，用于一起支付和发餐收餐
	UserName     string
	UserId       int    //下单用户的id
	Address      string //配送的地址
	ProductId    int    //下单的产品
	Num          int    //下单的数量
	UserPhone    string //下单用户的电话，从address数据库中获取
	Status       string //“未支付”“已支付”“送餐中”“已送达”“已评价”
	CanteenID    int    //订单属于哪个食堂
	DeliverID    int    //配送员编号
	DeliverName  string //配送员姓名
	DeliverPhone string //配送员电话
}
