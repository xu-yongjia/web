package api2

import (
	"gintest/DBstruct"

	"github.com/gin-gonic/gin"
)

type CarouselRequest struct {
	CanteenID   int    `json:"canteen_id"`
	CarouselURL string `json:"carousel_url"`
	Board       string `json:"board"`
}

func SaveCarouselURL(c *gin.Context) {

	var request CarouselRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(201, ERRRESPONSE("json参数格式有误", 201))
		return
	}

	var canteen DBstruct.Canteen
	if DBstruct.DB.Where("canteen_id = ?", request.CanteenID).First(&canteen).RecordNotFound() {
		c.JSON(201, ERRRESPONSE("餐厅不存在", 201))
		return
	}

	canteen.Board = request.Board
	canteen.ImgPath = request.CarouselURL
	if err := DBstruct.DB.Save(&canteen).Error; err != nil {
		c.JSON(201, ERRRESPONSE("数据库错误", 201))
		return
	}

	c.JSON(200, SUCCESSRESPONSE_NODATA())
}
