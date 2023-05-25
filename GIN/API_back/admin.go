//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 15:05:50
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:42:12
 */
package api2

// import (
// 	"cmall/pkg/logging"
// 	"cmall/service"
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// // AdminRegister 管理员注册接口
// func AdminRegister(c *gin.Context) {
// 	var service service.AdminRegisterService
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Register()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		fmt.Printf("err: %v\n", err)
// 		logging.Info(err)
// 	}
// }

// // AdminLogin 管理员登录接口
// func AdminLogin(c *gin.Context) {
// 	var service service.AdminLoginService
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Login()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }
