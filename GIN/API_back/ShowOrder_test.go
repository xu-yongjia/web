package api2

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestShowOrder(t *testing.T) {
	// 启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 2)

	url := "http://127.0.0.1:3000/api1/v2/admin/showOrder"
	// 配置请求内容
	data := make(map[string]interface{})                               
	data["page"] = 1                                     
	data["num_each_page"] = 10
	data["canteen_id"] = 1
	data["status"] = "送餐中"
	data["search_order_id"] = 1
	bytesData, err := json.Marshal(data)
	handleError(t, err)

	reqest, err := http.NewRequest("POST", url, bytes.NewReader(bytesData)) 
	handleError(t, err)
	reqest.Header.Add("Content-Type", "application/json")              
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
