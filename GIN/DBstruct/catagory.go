package DBstruct

type Category struct {
	CategoryID   uint `gorm:"primary_key"`
	CategoryName string
	CanteenID    uint `gorm:"primary_key;auto_increment:false"`
}
