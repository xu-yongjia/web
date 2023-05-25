package DBstruct

type ProductImg struct {
	ID        int    `json:"id" gorm:"primary_key"`
	ProductID int    `json:"product_id"` //菜品编号
	ImgPath   string `json:"img_path"`   //小图url，一道菜品可能有多条Product_img记录，用来存储多张图片的URL
}
