package DBstruct

// password不会被json传输，这样客户端在登录之后，不会有密码的缓存，相对安全（人走开后被偷看密码和修改密码）
type Adm struct {
	User_id     int    `json:"user_id" gorm:"primary_key"`
	UserName    string `json:"username" gorm:"unique"`
	Password    string `json:"-"`
	CanteenId   int
	CanteenName string
}
