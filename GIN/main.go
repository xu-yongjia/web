package main

import (
	api1 "gintest/API_front"
	"gintest/DBstruct"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() //加载环境变量
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	DBstruct.Database(dsn) //连接数据库
	api1.InitPay()         //初始化支付宝沙箱支付接口
	r := NewRouter()       //初始化路由
	r.Run(":3000")         //运行服务器
}
