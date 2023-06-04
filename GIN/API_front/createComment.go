package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	service := CreateCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(201, ERRRESPONSE(err.Error(), 201))
		// logging.Info(err)
	}
}

// CreateCommentService 评论创建的服务
type CreateCommentService struct {
	UserName       string `form:"user_name" json:"user_name"`
	ProductID      uint   `form:"product_id" json:"product_id"`
	ProductComment string `form:"product_comment" json:"product_comment"`
	Score          string `form:"score" json:"score"`
}

// Create 创建评论
func (service *CreateCommentService) Create() Response {
	comment := DBstruct.Comment{
		UserName:       service.UserName,
		ProductID:      service.ProductID,
		ProductComment: service.ProductComment,
		Score:          service.Score,
	}
	code := e.SUCCESS
	err := DBstruct.DB.Create(&comment).Error
	if err != nil {
		// logging.Info(err)
		code = e.ERROR_DATABASE
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//更新平均评分
	comments := []DBstruct.Comment{}
	total := 0
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
		if err := DBstruct.DB.Model(&product).Where("id=?", service.ProductID).Update("avg_score", avg).Error; err != nil {
			// logging.Info(err)
			code = e.ERROR_DATABASE
		}
	}
	return Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   BuildComment2(comment),
	}

}

// Comment 分类序列化器
type Comment_json2 struct {
	ID             uint   `json:"id"`
	ProductID      uint   `json:"ProductID"`
	UserName       string `json:"user_name"`
	ProductComment string `json:"ProductComment"`
	Score          string `json:"score"`
	CreatedAt      int64  `json:"commentDate"`
}

// BuildComment2 序列化分类
func BuildComment2(item DBstruct.Comment) Comment_json2 {
	return Comment_json2{
		ID:             item.ID,
		ProductID:      item.ProductID,
		UserName:       item.UserName,
		ProductComment: item.ProductComment,
		Score:          item.Score,
		CreatedAt:      item.CreatedAt.Unix(),
	}
}

// TrackedErrorResponse 有追踪信息的错误响应
