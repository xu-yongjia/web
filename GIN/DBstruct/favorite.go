package DBstruct

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	UserID   int
	ProducID int
}
