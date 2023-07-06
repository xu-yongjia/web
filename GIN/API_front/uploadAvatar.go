package api1

import (
	"fmt"
	"gintest/pkg/e"
	"mime"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// UploadToken 上传授权
func UploadToken(c *gin.Context) {
	service := UploadAvatarService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ERRRESPONSE(err.Error(), 201))
		//logging.Info(err)
	}
}

// UploadAvatarService 获得上传oss token的服务
type UploadAvatarService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadAvatarService) Post() Response {
	code := e.SUCCESS
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		// logging.Info(err)
		code = e.ERROR_OSS
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	fmt.Printf("bucket: %v\n", bucket) //debug
	if err != nil {
		// logging.Info(err)
		code = e.ERROR_OSS
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取扩展名
	ext := filepath.Ext(service.Filename)
	fmt.Printf("ext: %v\n", ext)

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	fmt.Printf("key: %v\n", key) //debug
	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	fmt.Printf("signedPutURL: %v\n", signedPutURL) //debug
	if err != nil {
		// logging.Info(err)
		code = e.ERROR_OSS
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 查看图片
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 24*3600*365*10)
	fmt.Printf("signedGetURL: %v\n", signedGetURL) //debug
	if err != nil {
		// logging.Info(err)
		code = e.ERROR_OSS
		return Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
