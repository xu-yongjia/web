package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// Category 分类序列化器
type Comment struct {
	ID             uint   `json:"id"`
	ProductID      uint   `json:"ProductID"`
	UserName       string `json:"user_name"`
	ProductComment string `json:"ProductComment"`
	Score          string `json:"score"`
	CreatedAt      int64  `json:"created_at"`
}

// ListComments 评论列表接口
func ListComments(c *gin.Context) {
	service := ListCommentsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ERRRESPONSE(err.Error(), 201))
	}
}

// BuildCategory 序列化分类
func BuildComment(item DBstruct.Comment) Comment {
	return Comment{
		ID:             item.ID,
		ProductID:      item.ProductID,
		UserName:       item.UserName,
		ProductComment: item.ProductComment,
		Score:          item.Score,
		CreatedAt:      item.CreatedAt.Unix(),
	}
}

// BuildCategories 序列化分类列表
func BuildComments(items []DBstruct.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}

// BuildListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) EResponse {
	return EResponse{
		Status: 200,
		Msg:    "ok",
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}

// ListCategoriesService 分类列表服务
type ListCommentsService struct {
	Limit     int  `form:"limit" json:"limit"`
	Start     int  `form:"start" json:"start"`
	ProductID uint `form:"productID" json:"productID"`
}

// List 视频列表
func (service *ListCommentsService) List() EResponse {

	comments := []DBstruct.Comment{}

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}
	if err := DBstruct.DB.Model(DBstruct.Comment{}).Where("product_id=?", service.ProductID).Count(&total).Error; err != nil {
		code = e.ERROR_DATABASE
		return EResponse{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err := DBstruct.DB.Where("product_id=?", service.ProductID).Limit(service.Limit).Offset(service.Start).Find(&comments).Error; err != nil {
		code = e.ERROR_DATABASE
		return EResponse{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return BuildListResponse(BuildComments(comments), uint(total))

}
