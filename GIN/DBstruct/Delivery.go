package DBstruct

import "github.com/jinzhu/gorm"

// Canteen 分类模型
type Delivery struct {
	gorm.Model
	Truename  string
	Phone     string
	CanteenID int
	Active    bool
}
