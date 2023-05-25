package DBstruct

import "github.com/jinzhu/gorm"

// Product 商品模型
type Product struct {
	gorm.Model
	//ID            int    `json:"product_id" gorm:"primary_key"`  //菜品编号
	Name          string  `json:"product_name"`                   //菜品名
	CanteenID     int     `json:"canteen_id"`                     //食堂编号
	CategoryID    int     `json:"category_id"`                    //类别ID
	Info          string  `json:"product_intro" gorm:"size:1000"` //菜品简介
	ImgPath       string  `json:"img_path"`                       //外面的大图
	Price         float32 `json:"product_selling_price"`          //原价
	DiscountPrice float32 `json:"product_price"`                  //折后价
	Title         string
}
