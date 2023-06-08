package api2

import (
	
	"os"
	"path/filepath"
	"mime"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Responsestruct represents the structure of the response that this service returns.
type Responsestruct struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

// UploadToken is a function to handle the request to upload an image.
func UploadToken(c *gin.Context) {
	service := UploadImgService{}
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, Responsestruct{
			Status: 201,  // Status code for request error
			Msg:    "json参数格式有误",
			Data:   map[string]string{},
		})
	}
}

// UploadImgService is a service that provides a method to upload an image.
type UploadImgService struct {
	Filename string `json:"filename"`
}

// Post is a method to handle the image upload.
func (service *UploadImgService) Post() Responsestruct {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return Responsestruct{
			Status: 404,
			Msg:    "token鉴权失败",
			Data:   map[string]string{},
		}
	}

	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return Responsestruct{
			Status: 404,
			Msg:    "token鉴权失败",
			Data:   map[string]string{},
		}
	}

	ext := filepath.Ext(service.Filename)

	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return Responsestruct{
			Status: 404,
			Msg:    "token鉴权失败",
			Data:   map[string]string{},
		}
	}

	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return Responsestruct{
			Status: 404,
			Msg:    "token鉴权失败",
			Data:   map[string]string{},
		}
	}

	return Responsestruct{
		Status: 200,
		Msg:    "OK",
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
