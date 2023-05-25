package DBstruct

import "github.com/jinzhu/gorm"

type Carousel struct { //轮播图
	gorm.Model
	ImgPath   string
	ProductId uint //对应的菜品ID
}
