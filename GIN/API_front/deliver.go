//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:40:36
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 20:58:01
 */
package api1

// import (
// 	"cmall/pkg/logging"
// 	"cmall/service"
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// // CreateProduct 创建商品
// func CreateDeliver(c *gin.Context) {
// 	service := service.CreateDeliverService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Create()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }

// // ListProducts 商品列表接口
// func ListDelivers(c *gin.Context) {
// 	var service service.ListDeliversService
// 	if err := c.ShouldBind(&service); err == nil {
// 		fmt.Printf("service: %v\n", service)
// 		res := service.List()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}

// }

// // // ShowProduct 商品详情接口
// // func ShowDeliver(c *gin.Context) {
// // 	service := service.ShowDeliverService{}
// // 	res := service.Show(c.Param("id"))
// // 	c.JSON(200, res)
// // }

// // DeleteProduct 删除商品的接口
// func DeleteDeliver(c *gin.Context) {
// 	service := service.DeleteDeliverService{}
// 	res := service.Delete(c.Param("id"))
// 	c.JSON(200, res)
// }

// // UpdateProduct 更新商品的接口
// func UpdateDeliver(c *gin.Context) {
// 	service := service.UpdateDeliverService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Update()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }

// // // SearchDelivers 搜索配送员的接口
// // func SearchDelivers(c *gin.Context) {
// // 	service := service.SearchDeliversService{}
// // 	if err := c.ShouldBind(&service); err == nil {
// // 		res := service.Show()
// // 		c.JSON(200, res)
// // 	} else {
// // 		c.JSON(200, ErrorResponse(err))
// // 		logging.Info(err)
// // 	}
// // }

// func DeliverOrder(c *gin.Context) {
// 	service := service.DeliverOrderService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Deliver()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }
