package api

import (
	"bytes"
	"encoding/json"
	"gintest/DBstruct"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}
func server() {
	//关闭gin的日志输出
	gin.SetMode(gin.ReleaseMode)
	//数据库的日志输出在DBStruct/init.go中使用语句db.LogMode()配置，已改为false
	//打开数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	DBstruct.Database(dsn)
	//在这里配置一个路由器并运行，配置的路由器时不加token中间件，这样发送的请求不需要带token
	r := gin.Default()
	v1 := r.Group("/api/")
	v1.POST("GetCategoryList", GetCategoryList)
	v1.POST("AddCategory", AddCategory)
	v1.POST("UpdateCategory", UpdateCategory)
	v1.POST("HasDelCategory", HasDelCategory)
	v1.POST("DelCategory", DelCategory)
	r.Run(":3000")
}

func TestGetCategoryList(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/api/GetCategoryList"
	//配置请求内容
	data := make(map[string]interface{})                                  //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["canteen_id"] = 1                                                //配置data的内容
	bytesData, _ := json.Marshal(data)                                    //将data编码为json格式，byte[]类型
	reqest, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData)) //用url和bytesData创建一个请求结构体
	reqest.Header.Add("Content-Type", "application/json")                 //设置头部字段Content-Type为json格式
	//发送请求，获取返回数据
	client := &http.Client{}
	resp, _ := client.Do(reqest)
	body, _ := ioutil.ReadAll(resp.Body)
	//格式化返回的json数据并打印到日志
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	t.Log(str.String())
}

func TestAddCategory(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/api/AddCategory"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["canteen_id"] = 1
	data["category_name"] = "烧卤"                                          //配置data的内容
	bytesData, _ := json.Marshal(data)                                    //将data编码为json格式，byte[]类型
	reqest, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData)) //用url和bytesData创建一个请求结构体
	reqest.Header.Add("Content-Type", "application/json")                 //设置头部字段Content-Type为json格式
	//发送请求，获取返回数据
	client := &http.Client{}
	resp, _ := client.Do(reqest)
	body, _ := ioutil.ReadAll(resp.Body)
	//格式化返回的json数据并打印到日志
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	t.Log(str.String())
}

func TestUpdateCategory(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/api/UpdateCategory"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["category_id"] = 1
	data["category_name"] = "自选"                                          //配置data的内容
	bytesData, _ := json.Marshal(data)                                    //将data编码为json格式，byte[]类型
	reqest, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData)) //用url和bytesData创建一个请求结构体
	reqest.Header.Add("Content-Type", "application/json")                 //设置头部字段Content-Type为json格式
	//发送请求，获取返回数据
	client := &http.Client{}
	resp, _ := client.Do(reqest)
	body, _ := ioutil.ReadAll(resp.Body)
	//格式化返回的json数据并打印到日志
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	t.Log(str.String())
}

func TestHasDelCategory(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/api/HasDelCategory"
	//配置请求内容
	data := make(map[string]interface{})                                  //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["category_id"] = 123                                             //配置data的内容
	bytesData, _ := json.Marshal(data)                                    //将data编码为json格式，byte[]类型
	reqest, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData)) //用url和bytesData创建一个请求结构体
	reqest.Header.Add("Content-Type", "application/json")                 //设置头部字段Content-Type为json格式
	//发送请求，获取返回数据
	client := &http.Client{}
	resp, _ := client.Do(reqest)
	body, _ := ioutil.ReadAll(resp.Body)
	//格式化返回的json数据并打印到日志
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	t.Log(str.String())
}

func TestDelCategory(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/api/DelCategory"
	//配置请求内容
	data := make(map[string]interface{})                                  //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["category_id"] = 123                                             //配置data的内容
	bytesData, _ := json.Marshal(data)                                    //将data编码为json格式，byte[]类型
	reqest, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData)) //用url和bytesData创建一个请求结构体
	reqest.Header.Add("Content-Type", "application/json")                 //设置头部字段Content-Type为json格式
	//发送请求，获取返回数据
	client := &http.Client{}
	resp, _ := client.Do(reqest)
	body, _ := ioutil.ReadAll(resp.Body)
	//格式化返回的json数据并打印到日志
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	t.Log(str.String())
}
