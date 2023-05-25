package DBstruct

// password不会被json传输，这样客户端在登录之后，不会有密码的缓存，相对安全（人走开后被偷看密码和修改密码）
type User struct {
	Phone    string `json:"phone"` //这里有一个“电话”信息，但是系统利用的电话信息主要是从Address记录中获取的，与这里的无关，只是为了保证扩展性而添加的冗余字段
	User_id  int    `json:"user_id" gorm:"primary_key"`
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}
