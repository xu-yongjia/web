package api2

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type PaginationRequest struct {
	Page        int `json:"page"`
	NumEachPage int `json:"num_each_page"`
}

type AddressDisplay struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"` // Add this line
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type UserDisplay struct {
	ID        uint             `json:"user_id"`
	UserName  string           `json:"user_name"`
	Phone     string           `json:"phone"`
	Addresses []AddressDisplay `json:"addresss"`
}

func ShowUser(c *gin.Context) {

	var pagination PaginationRequest
	if err := c.ShouldBindJSON(&pagination); err != nil {
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
		return
	}

	var users []DBstruct.User
	DBstruct.DB.Limit(pagination.NumEachPage).Offset((pagination.Page - 1) * pagination.NumEachPage).Find(&users)

	var userDisplays []UserDisplay
	for _, user := range users {
		var userDisplay UserDisplay
		userDisplay.ID = user.ID
		userDisplay.UserName = user.UserName
		userDisplay.Phone = user.Phone

		var addresses []DBstruct.Address
		DBstruct.DB.Where("user_id = ?", user.ID).Find(&addresses)
		for _, address := range addresses {
			var addressDisplay AddressDisplay
			addressDisplay.ID = address.ID
			addressDisplay.UserID = user.ID
			addressDisplay.UserName = address.UserName
			addressDisplay.Phone = address.Phone
			addressDisplay.Address = address.Address
			userDisplay.Addresses = append(userDisplay.Addresses, addressDisplay)
		}

		userDisplays = append(userDisplays, userDisplay)
	}

	var totalCount int
	DBstruct.DB.Model(&DBstruct.User{}).Count(&totalCount)

	c.JSON(200, SUCCESSRESPONSE(gin.H{
		"count":    totalCount,
		"userlist": userDisplays,
	}))
}
