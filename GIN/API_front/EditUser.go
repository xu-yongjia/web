package api1

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type EditUserRequest struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func EditAccount(c *gin.Context) {
	var request EditUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(201, ERRRESPONSE("json参数格式有误", 201))
		return
	}

	var user DBstruct.User
	if DBstruct.DB.Where("id = ?", request.UserID).First(&user).RecordNotFound() {
		c.JSON(201, ERRRESPONSE("用户不存在", 201))
		return
	}

	user.UserName = request.UserName
	user.Password = request.Password
	if len(request.Avatar) > 0 {
		user.Avatar = request.Avatar
	}
	if err := DBstruct.DB.Save(&user).Error; err != nil {
		c.JSON(201, ERRRESPONSE("数据库错误", 201))
		return
	}

	c.JSON(200, SUCCESSRESPONSE(BuildUser(user)))
}
