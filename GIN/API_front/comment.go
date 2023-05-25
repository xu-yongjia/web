package api1

// import (
// 	"cmall/pkg/logging"
// 	"cmall/service"

// 	"github.com/gin-gonic/gin"
// )

// // ListCategories 分类列表接口
// func ListComments(c *gin.Context) {
// 	service := service.ListCommentsService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.List()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }
