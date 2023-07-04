package api2

import (
	"bytes"
	"encoding/json"
	//"gintest/DBstruct"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	//"github.com/gin-gonic/gin"
)




func TestAdminLogin(t *testing.T) {
	// 启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 2)  // 添加这一行

	url := "http://127.0.0.1:3000/api1/v2/admin/login"
	// 配置请求内容
	data := make(map[string]interface{})                               // 创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["canteen_id"] = 1                                             // 配置data的内容
	data["password"] = "123"                                         // 设置密码
	bytesData, err := json.Marshal(data)
	handleError(t, err)

	reqest, err := http.NewRequest("POST", url, bytes.NewReader(bytesData)) // 用url和bytesData创建一个请求结构体
	handleError(t, err)
	reqest.Header.Add("Content-Type", "application/json")              // 设置头部字段Content-Type为json格式
	// 发送请求，获取返回数据
	client := &http.Client{}
	resp, err := client.Do(reqest)
	handleError(t, err)
	defer resp.Body.Close()
	
	if resp == nil {
		t.Fatal("Received nil response")
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	// 格式化返回的json数据并打印到日志
	var str bytes.Buffer
	err = json.Indent(&str, []byte(body), "", "    ")
	handleError(t, err)

	t.Log(str.String())
}
