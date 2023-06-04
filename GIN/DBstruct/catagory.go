package DBstruct

type Category struct {
	CategoryID   uint
	CategoryName string
	CanteenID    uint `gorm:"primary_key"`
}
