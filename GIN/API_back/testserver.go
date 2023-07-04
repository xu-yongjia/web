package api2

import (
	
	"gintest/DBstruct"
	
	"testing"
	
	"github.com/gin-gonic/gin"
)

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func server() {
    // 关闭gin的日志输出
    gin.SetMode(gin.ReleaseMode)
    // 数据库的日志输出在DBStruct/init.go中使用语句db.LogMode()配置，已改为false
    // 打开数据库
    dsn := "root:Aa15606936638@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
    DBstruct.Database(dsn)
    // 在这里配置一个路由器并运行，配置的路由器时不加token中间件，这样发送的请求不需要带token
    r := gin.Default()
    v1 := r.Group("/api1/v2")
    v1.POST("admin/login", AdminLogin)
	v1.POST("admin/assignDelivery",AssignDelivery)
	v1.POST("admin/saveCarouselUrl",SaveCarouselURL)
	v1.POST("admin/showDelivery",ShowDelivery)
	v1.POST("admin/showOrder",ShowOrder)
	v1.POST("admin/showUser",ShowUser)
    r.Run(":3000")
}

