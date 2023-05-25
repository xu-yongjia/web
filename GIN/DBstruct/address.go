package DBstruct

// Address 收货地址模型
type Address struct {
	ID       int    `json:"id" gorm:"primary_key"`
	User_id  int    `json:"user_id"`
	UserName string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
