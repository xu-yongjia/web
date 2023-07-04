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
	gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	DBstruct.Database(dsn)
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
	go server()

	url := "http://127.0.0.1:3000/api/GetCategoryList"

	data := make(map[string]interface{})
	data["canteen_id"] = 1
	bytesData, _ := json.Marshal(data)
	reqest, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	reqest.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(reqest)
	body, _ := ioutil.ReadAll(resp.Body)

	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	t.Log(str.String())
}
