package DBstruct

// Canteen 分类模型
type Delivery struct {
	ID            int `gorm:"primary_key,unique"`
	Truename      string
	Phone         string
	CanteenNumber int
	Active        bool
}
