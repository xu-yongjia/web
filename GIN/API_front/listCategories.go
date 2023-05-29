package api1

import (
	"gintest/DBstruct"
	"time"

	"github.com/gin-gonic/gin"
)

type categoryJson struct {
	CategoryID   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func getJsonCategory(src DBstruct.Category) categoryJson {
	return categoryJson{
		CategoryID:   src.CategoryID,
		CategoryName: src.CategoryName,
		CreatedAt:    src.CreatedAt,
	}
}

// 读取报文体的palceID
func ListCategories(c *gin.Context) {
	type msg struct {
		CanteenID uint `form:"placeID" json:"placeID"`
	}
	var m msg
	if e := c.ShouldBindJSON(&m); e == nil {
		categoryRecords := make([]DBstruct.Category, 0)
		if e = DBstruct.DB.Where("canteen_id = ?", m.CanteenID).Find(&categoryRecords).Error; e == nil {
			categoryJsons := make([]categoryJson, 0)
			for _, categoryRecord := range categoryRecords {
				categoryJson := getJsonCategory(categoryRecord)
				categoryJsons = append(categoryJsons, categoryJson)
			}
			type returnmsg struct {
				CategoryJsons []categoryJson `json:"items"`
			}
			c.JSON(200, SUCCESSRESPONSE(returnmsg{CategoryJsons: categoryJsons}))
		} else {
			c.JSON(201, ERRRESPONSE(e.Error(), 201))
		}
	} else {
		c.JSON(400, ERRRESPONSE("数据格式错误", 201))
	}
}
