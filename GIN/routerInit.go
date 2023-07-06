package main

import (
	api2 "gintest/API_back"
	"gintest/API_back/api"
	api1 "gintest/API_front"
	"gintest/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")
	{
		// 商品详情
		v1.POST("product/getDetails", api1.ShowProduct)
		// 获取菜品图片url
		v1.POST("product/getDetailsPicture", api1.ShowProductImgs)
		// 获取评论
		v1.POST("product/getDetailsComment", api1.ListComments)
		// 获取评分排行榜
		v1.POST("rankings", api1.ListRanking)
		// 获取有哪些菜品类别
		v1.POST("categories", api1.ListCategories)
		// 获取菜品列表
		v1.POST("products", api1.ListProducts)
		// 轮播图
		v1.GET("carousels", api1.ListCarousels)

		//沙箱支付回调函数1
		v1.GET("/alipay/callback", api1.Callback)
		//沙箱支付回调函数2
		v1.POST("/alipay/notify", api1.Notify)
		//用户注册
		v1.POST("users/register", api1.RegistUser)
		//用户登录
		v1.POST("users/login", api1.UserLogin)

		authed := v1.Group("users/")
		authed.Use(middleware.JWT())
		{
			//获取信息数量统计
			authed.POST("showCount", api1.ShowCount)
			//上传头像
			authed.POST("uploadAvatar", api1.UploadToken)
			//更改用户名
			authed.POST("editAccount", api1.EditAccount)
			//添加收藏
			authed.POST("collect/addCollect", api1.AddCollect)
			//支付订单
			authed.POST("payments", api1.Pay)
			//下单
			authed.POST("createOrder", api1.CreateOrder)
			//确认收餐
			authed.POST("recieveOrder", api1.RecieveOrder)
			// 获取购物车
			authed.POST("getCart", api1.GetCart)
			// 获取订单
			authed.POST("getorder", api1.GetOrder)
			// 		//验证token
			// 		authed.GET("ping", api1.CheckToken)
			// 		//用户操作
			// 		authed.PUT("user", api1.UserUpdate)
			// 		authed.POST("user/sending-email", api1.SendEmail)
			// 		// 上传操作
			// 		authed.POST("avatar", api1.UploadToken)
			// 		//收藏夹操作
			// 提交评论
			authed.POST("submitComment", api1.CreateComment)

			authed.POST("collect/getCollect", api1.GetCollect)

			authed.POST("shoppingCart/addShoppingCart", api1.AddShoppingCart)
			authed.POST("shoppingCart/updateShoppingCart", api1.UpdateShoppingCart)
			authed.POST("shoppingCart/deleteShoppingCart", api1.DeleteShoppingCart)

			authed.POST("getUseraddress", api1.GetUserAddress)
			authed.POST("postUseraddress", api1.PostUseraddress)
			authed.POST("deleteAddress", api1.DeleteAddress)
			authed.POST("saveEdit", api1.EditAddress)
		}
	}
	v2 := r.Group("/api1/v2")
	{
		v2.POST("admin/login", api2.AdminLogin)

		authed2 := v2.Group("/")
		// 	//登录验证
		authed2.Use(middleware.JWT())
		{ // 上传操作
			authed2.POST("admin/generateurl", api2.UploadToken)
			// 	// 查看用户信息
			authed2.POST("admin/showUser", api2.ShowUser)
			// 	// 查看订单信息
			authed2.POST("admin/showOrder", api2.ShowOrder)
			// 	// 指派配送员
			authed2.POST("admin/assignDelivery", api2.AssignDelivery)
			// 请求图片空间
			authed2.POST("admin/uploadImg", api2.UploadToken)
			//保存轮播图片路径
			authed2.POST("admin/saveCarouselUrl", api2.SaveCarouselURL)
			//显示配送员
			authed2.POST("admin/showDelivery", api2.ShowDelivery)

			authed2.POST("admin/delivery/add", api.AddDelivery)
			authed2.POST("admin/delivery/update", api.UpdateDelivery)
			authed2.POST("admin/delivery/del", api.DelDelivery)

			authed2.POST("admin/category/list", api.GetCategoryList)
			authed2.POST("admin/category/add", api.AddCategory)
			authed2.POST("admin/category/update", api.UpdateCategory)
			authed2.POST("admin/category/hasDel", api.HasDelCategory)
			authed2.POST("admin/category/del", api.DelCategory)

			authed2.POST("admin/product/list", api.GetProductList)
			authed2.POST("admin/product/add", api.AddProduct)
			authed2.POST("admin/product/update", api.UpdateProduct)
			authed2.POST("admin/product/del", api.DelProduct)
			authed2.POST("admin/product/photo/add", api.SavePhoto)
		}
	}
	return r
}
