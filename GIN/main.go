package main

import (
	api1 "gintest/API_front"
	"gintest/DBstruct"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	DBstruct.Database(dsn)
	api1.InitPay()
	r := NewRouter()
	r.Run(":3000")
	// token, _ := util.GenerateToken("123321", "321123", 2)
	// result, _ := util.ParseToken(token)
	// fmt.Print(result)
}
