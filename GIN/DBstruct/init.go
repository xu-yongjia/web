package DBstruct

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(false)
	// Error
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//默认不加复数
	db.SingularTable(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&Address{}).
		AutoMigrate(&Canteen{}).
		AutoMigrate(&Carousel{}).
		AutoMigrate(&Cart{}).
		AutoMigrate(&Category{}).
		AutoMigrate(&Delivery{}).
		AutoMigrate(&Favorite{}).
		AutoMigrate(&Order{}).
		AutoMigrate(&ProductImg{}).
		AutoMigrate(&Product{}).
		AutoMigrate(&User{}).
		AutoMigrate(&Canteen{}).
		AutoMigrate(&Comment{})
}
