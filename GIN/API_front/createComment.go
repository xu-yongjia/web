package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	type msg struct {
		UserID         uint   `form:"user_id" json:"user_id"`
		ProductID      uint   `form:"product_id" json:"product_id"`
		ProductComment string `form:"product_comment" json:"product_comment"`
		Score          string `form:"score" json:"score"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		var user DBstruct.User
		if e = DBstruct.DB.Where("id = ?", m.UserID).First(&user).Error; e == nil { //按照CanteenID查找账号
			service := CreateCommentService{
				UserID:         int(m.UserID),
				UserName:       user.UserName,
				ProductID:      m.ProductID,
				ProductComment: m.ProductComment,
				Score:          m.Score,
			}
			res := service.Create()
			c.JSON(200, res)
		} else {
			c.JSON(201, ERRRESPONSE("账号不存在"+e.Error(), 201)) //具有请求中的CanteenID的账号不存在
		}
	} else { //JSON解析失败
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}

// CreateCommentService 评论创建的服务
type CreateCommentService struct {
	OrderID        uint   `form:"order_id" json:"order_id"`
	UserID         int    `form:"user_id" json:"user_id"`
	UserName       string `form:"user_name" json:"user_name"`
	ProductID      uint   `form:"product_id" json:"product_id"`
	ProductComment string `form:"product_comment" json:"product_comment"`
	Score          string `form:"score" json:"score"`
}

// Create 创建评论
func (service *CreateCommentService) Create() Response {
	comment := DBstruct.Comment{
		UserID:         service.UserID,
		UserName:       service.UserName,
		ProductID:      service.ProductID,
		ProductComment: service.ProductComment,
		Score:          service.Score,
	}
	var historyComment DBstruct.Comment
	err := DBstruct.DB.Where("user_id=? and product_id=?", service.UserID, service.ProductID).First(&historyComment).Error //查找该用户是否对该商品进行过评价，如有则覆盖
	if err == nil {
		historyComment.UserName = service.UserName
		historyComment.ProductComment = service.ProductComment
		historyComment.Score = service.Score
		err = DBstruct.DB.Save(&historyComment).Error
		if err != nil {
			return Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
			}
		}
	} else { //如没有则新建
		code := e.SUCCESS
		err := DBstruct.DB.Create(&comment).Error
		if err != nil {
			code = e.ERROR_DATABASE
			return Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	//更新order记录的状态
	DBstruct.DB.Model(&DBstruct.Order{}).Where("order_id = ? AND product_id = ?", service.OrderID, service.ProductID).Update("status", "已评价")
	//更新平均评分
	comments := []DBstruct.Comment{}
	total := 0
	code := e.SUCCESS
	if err := DBstruct.DB.Model(DBstruct.Comment{}).Where("product_id=?", service.ProductID).Count(&total).Error; err != nil {
		// logging.Info(err)
		code = e.ERROR_DATABASE
	}
	if total != 0 {
		if err := DBstruct.DB.Where("product_id=?", service.ProductID).Find(&comments).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
		}

		var sum float64
		// strconv.FormatFloat(3.1415926, 'f', 5, 64) // 3.14159 小数变字符串
		// strconv.ParseFloat("123.213", 64)          // 123.213 <nil> 字符串变小数
		for _, v := range comments {
			s, _ := strconv.ParseFloat(v.Score, 64)
			sum = sum + s
		}
		avg := strconv.FormatFloat(sum/float64(total), 'f', 2, 64)

		var product DBstruct.Product
		if err := DBstruct.DB.Model(&product).Where("id=?", service.ProductID).Update("score", avg).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
		}
	}
	return Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}

// TrackedErrorResponse 有追踪信息的错误响应
