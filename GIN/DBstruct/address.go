package DBstruct

import "github.com/jinzhu/gorm"

// Address 收货地址模型
type Address struct {
	gorm.Model
	User_id  int    `json:"user_id"`
	UserName string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
