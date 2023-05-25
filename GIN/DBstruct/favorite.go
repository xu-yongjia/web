package DBstruct

type Favorite struct {
	Record_id int `gorm:"primary_key"`
	UserID    int
	ProducID  int
}
