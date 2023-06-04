package DBstruct

import "github.com/jinzhu/gorm"

// Canteen 分类模型
type Canteen struct {
	gorm.Model
	CanteenID   int
	CanteenName string
	ImgPath     string
	Password    string
}
