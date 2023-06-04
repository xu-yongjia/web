package DBstruct

import "github.com/jinzhu/gorm"

type Carousel struct { //轮播图
	gorm.Model
	ImgPath   string
	CanteenId uint //对应的食堂ID
}
