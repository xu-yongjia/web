package main

import (
	api2 "gintest/API_back"
	api1 "gintest/API_front"
	"gintest/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(middleware.Cors())
	// store := cookie.NewStore([]byte(sdk.VERSION))
	// r.Use(sessions.Sessions("mysession", store))
	// 路由

	v1 := r.Group("/api/v1/")
	{
		//商品详情
		v1.POST("product/getDetails", api1.ShowProduct)
		//获取菜品图片url
		v1.POST("product/getDetailspicture", api1.ShowProductImgs)
		//获取评论
		v1.POST("product/getDetailsComment", api1.ListComments)
		//获取评分排行榜
		v1.POST("rankings", api1.ListRanking)
		//获取有哪些菜品类别
		v1.POST("categories", api1.ListCategories)
		//获取菜品列表
		v1.POST("products", api1.ListProducts)
		//轮播图
		v1.GET("carousels", api1.ListCarousels)
		// 	// 用户注册
		v1.POST("users/register", api1.RegistUser)
		//	v1.POST("user/register", api1.UserRegister)
		// 	// 用户登录
		v1.POST("users/login", api1.UserLogin)
		// 	//商品操作
		// 	v1.GET("products", api1.ListProducts)
		// 	v1.GET("products/:id", api1.ShowProduct)
		// 	//轮播图操作
		// 	v1.GET("carousels", api1.ListCarousels)
		// 	//商品图片操作
		// 	v1.GET("imgs/:id", api1.ShowProductImgs)
		//商品详情图片操作
		//v1.GET("info-imgs/:id", api.ShowInfoImgs)
		// 	//商品参数图片操作
		// 	v1.GET("param-imgs/:id", api1.ShowParamImgs)
		// 	//分类操作
		// 	v1.GET("categories", api1.ListCategories)
		// 	//搜索操作
		// 	v1.POST("searches", api1.SearchProducts)
		// 	//排行榜/热门
		// 	v1.GET("rankings", api1.ListRanking)
		// 	v1.GET("elec-rankings", api1.ListElecRanking)
		// 	v1.GET("acce-rankings", api1.ListAcceRanking)
		// 	//README操作
		// 	v1.GET("notices", api1.ShowNotice)
		// 	v1.GET("geetest", api1.InitGeetest)
		// 	//支付操作
		// 	v1.GET("payments", api1.ConfirmPay)
		// 	// 需要登录保护的

		// s.GET("/alipay/pay", pay)
		v1.GET("/alipay/callback", api1.Callback)
		v1.POST("/alipay/notify", api1.Notify)

		// s.Run(":" + kServerPort)
		authed := v1.Group("users/")
		authed.Use(middleware.JWT())
		{
			authed.POST("payments", api1.Pay)
			authed.POST("createOrder", api1.CreateOrder)
			//获取购物车
			authed.POST("getCart", api1.GetCart)
			//获取订单
			authed.POST("getorder", api1.GetOrder)
			// 		//验证token
			// 		authed.GET("ping", api1.CheckToken)
			// 		//用户操作
			// 		authed.PUT("user", api1.UserUpdate)
			// 		authed.POST("user/sending-email", api1.SendEmail)
			// 		// 上传操作
			// 		authed.POST("avatar", api1.UploadToken)
			// 		//收藏夹操作
			//提交评论
			authed.POST("comment", api1.CreateComment)
			authed.POST("collect/addCollect", api1.AddCollect)
			// 		authed.GET("favorites/:id", api1.ShowFavorites)
			// 		authed.POST("favorites", api1.CreateFavorite)
			// 		authed.DELETE("favorites", api1.DeleteFavorite)
			// 		//订单操作
			// 		authed.POST("orders", api1.CreateOrder)
			// 		authed.GET("user/:id/orders", api1.ListOrders)
			// 		authed.GET("orders/:num", api1.ShowOrder)
			// 		//购物车操作
			authed.POST("shoppingCart/addShoppingCart", api1.AddShoppingCart)
			authed.POST("shoppingCart/updateShoppingCart", api1.UpdateShoppingCart)
			authed.POST("shoppingCart/deleteShoppingCart", api1.DeleteShoppingCart)
			// 		authed.POST("carts", api1.CreateCart)
			// 		authed.GET("carts/:id", api1.ShowCarts)
			// 		authed.PUT("carts", api1.UpdateCart)
			// 		authed.DELETE("carts", api1.DeleteCart)
			// 		//收货地址操作
			// 		authed.POST("addresses", api1.CreateAddress)
			authed.POST("getUseraddress", api1.GetUserAddress)
			authed.POST("postUseraddress", api1.PostUseraddress)
			authed.POST("deleteAddress", api1.DeleteAddress)
			authed.POST("saveEdit", api1.EditAddress)
			// 		authed.PUT("addresses", api1.UpdateAddress)
			// 		authed.DELETE("addresses", api1.DeleteAddress)
			// 		//支付操作
			// 		authed.POST("payments", api1.InitPay)

			// 		//数量操作
			// 		authed.GET("counts/:id", api1.ShowCount)

		}

	}
	v2 := r.Group("/api1/v2")
	{
		// 	// 管理员注册
		// 	v2.POST("admin/register", api2.AdminRegister)
		// 	// 管理员登录
		v2.POST("admin/login", api2.AdminLogin)

		// 	//商品操作
		//	v2.GET("products", api2.ListProducts)
		// 	v2.GET("products/:id", api2.ShowProduct)
		// 	//轮播图操作
		// 	v2.GET("carousels", api2.ListCarousels)
		// 	//商品图片操作
		// 	v2.GET("imgs/:id", api2.ShowProductImgs)
		// 	//分类操作
		// 	v2.GET("categories", api2.ListCategories)
		authed2 := v2.Group("/")
		// 	//登录验证
		authed2.Use(middleware.JWT())
		{ // 上传操作
			authed2.POST("avatar", api2.UploadToken)
			// 	// 查看用户信息
			authed2.POST("admin/showUser", api2.ShowUser)
			// 	// 查看订单信息
			authed2.POST("admin/showOrder", api2.ShowOrder)
			// 	// 指派配送员
			authed2.POST("admin/assignDelivery", api2.AssignDelivery)
			// 请求图片空间
			 authed2.POST("admin/uploadImg", api2.UploadToken)
			 //保存轮播图片路径
			 authed2.POST("admin/save_carousel_url", api2.SaveCarouselURL)

			// 		//商品操作
			// 		authed2.POST("products", api2.CreateProduct)
			// 		authed2.DELETE("products/:id", api2.DeleteProduct)
			// 		authed2.PUT("products", api2.UpdateProduct)
			// 		//轮播图操作
			// 		authed2.POST("carousels", api2.CreateCarousel)
			// 		//商品图片操作
			// 		authed2.POST("imgs", api2.CreateProductImg)
			// 		//商品详情图片操作
			// 		authed2.POST("info-imgs", api2.CreateInfoImg)
			// 		//商品参数图片操作
			// 		authed2.POST("param-imgs", api2.CreateParamImg)
			// 		//分类操作
			// 		authed2.POST("categories", api2.CreateCategory)
			// 		//公告操作
			// 		authed2.POST("notices", api2.CreateNotice)
			// 		authed2.PUT("notices", api2.UpdateNotice)
		}
	}
	return r
}
