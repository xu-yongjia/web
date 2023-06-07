package api1

import (
	"gintest/DBstruct"
	"gintest/pkg/e"

	"github.com/gin-gonic/gin"
)

// UserUpdate 用户修改信息接口
func UserUpdate(c *gin.Context) {
	var service UserUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)

	} else {
		c.JSON(200, ERRRESPONSE(err.Error(), 201))
	}
}

// UserUpdateService 用户修改信息的服务
type UserUpdateService struct {
	ID       uint   `form:"id" json:"id"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// Update 用户修改信息
func (service *UserUpdateService) Update() Response {
	var user DBstruct.User
	code := e.SUCCESS
	//找到用户
	err := DBstruct.DB.First(&user, service.ID).Error
	if err != nil {

		code = e.ERROR_DATABASE
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	user.UserName = service.UserName
	if service.Avatar != "" {
		user.Avatar = service.Avatar
	}
	err = DBstruct.DB.Save(&user).Error
	if err != nil {

		code = e.ERROR_DATABASE
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: User{
			ID:        user.ID,
			UserName:  user.UserName,
			Avatar:    user.AvatarURL(),
			CreatedAt: user.CreatedAt.Unix(),
		},
	}
}

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
