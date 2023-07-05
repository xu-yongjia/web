package api1

import (
	"bytes"
	"encoding/json"
	"gintest/DBstruct"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

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
	//打开数据库
	dsn := "root:Aa15606936638@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	DBstruct.Database(dsn)
	//在这里配置一个路由器并运行，配置的路由器不使用token中间件，这样发送的请求不需要带token
	r := gin.Default()
	v1 := r.Group("/apifront")
	v1.POST("product/getDetails", ShowProduct)
	v1.POST("product/getDetailspicture", ShowProductImgs)
	v1.POST("product/getDetailsComment", ListComments)
	v1.POST("rankings", ListRanking) //这个没写，看不懂要请求什么消息
	v1.POST("categories", ListCategories)
	v1.POST("products", ListProducts)
	v1.GET("carousels", ListCarousels) 
	v1.POST("users/register", RegistUser)
	v1.POST("users/login", UserLogin)

	v1.GET("/alipay/callback", Callback) //我感觉这个和notify都是不用写的，你到时看看，如果要写把data补上即可
	v1.POST("/alipay/notify", Notify)

	authed := v1.Group("users/")
	//authed.Use(middleware.JWT())//这段应该不要不然就要用token验证了
	{
		authed.POST("payments", Pay) 
		authed.POST("createOrder", CreateOrder)
		authed.POST("getCart", GetCart)
		authed.POST("getorder", GetOrder)
		authed.POST("comment", CreateComment)
		authed.POST("collect/addCollect", AddCollect)
		authed.POST("collect/getCollect", GetCollect)

		authed.POST("shoppingCart/addShoppingCart", AddShoppingCart)
		authed.POST("shoppingCart/updateShoppingCart", UpdateShoppingCart)
		authed.POST("shoppingCart/deleteShoppingCart", DeleteShoppingCart)

		authed.POST("getUseraddress", GetUserAddress)
		authed.POST("postUseraddress", PostUseraddress)
		authed.POST("deleteAddress", DeleteAddress)
		authed.POST("saveEdit", EditAddress)
		r.Run(":3000")
	}
}
func TestShowProduct(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/product/getDetails"

	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["productID"] = 1                //配置data的内容
	bytesData, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	reqest, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		t.Fatal(err)
	} //用url和bytesData创建一个请求结构体
	reqest.Header.Add("Content-Type", "application/json") //设置头部字段Content-Type为json格式
	//发送请求，获取返回数据
	client := &http.Client{}
	resp, _ := client.Do(reqest)
	body, _ := ioutil.ReadAll(resp.Body)
	//格式化返回的json数据并打印到日志
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")
	t.Log(str.String())
}

func TestShowProductImgs(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/product/getDetailspicture"
	//配置请求内容
	data := make(map[string]interface{})                                  //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["productID"] = 1                                                 //配置data的内容
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
func TestListComments(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/product/getDetailsComment"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["limit"] = 1
	data["start"] = 2
	data["productID"] = 3                                                 //配置data的内容
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

func TestListCategories(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/categories"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["limit"] = 1
	data["start"] = 2
	data["placeID"] = 3                                                   //配置data的内容
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

func TestListProducts(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/products"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["limit"] = 1
	data["start"] = 2
	data["placeID"] = 3
	data["categoryID"] = 4                                                //配置data的内容
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

func TestListCarousels(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/carousels"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["id"] = 1
	data["img_path"] = "http://..."
	data["product_id"] = 3
	data["created_at"] = 4                                                //配置data的内容
	bytesData, _ := json.Marshal(data)                                    //将data编码为json格式，byte[]类型
	reqest, _ := http.NewRequest("GET", url, bytes.NewReader(bytesData)) //用url和bytesData创建一个请求结构体
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

func TestRegistUser(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/users/register"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["userName"] = "小明"
	data["password"] = "12345678"                                         //配置data的内容
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
func TestUserLogin(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/users/login"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["username"] = "小明"
	data["user_id"] = "1"
	data["phone"] = "12345678910"
	data["-"] = "12345678"                                                //配置data的内容
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

func TestCallback(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/alipay/callback"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	//配置data的内容
	bytesData, _ := json.Marshal(data)                                    //将data编码为json格式，byte[]类型
	reqest, _ := http.NewRequest("GET", url, bytes.NewReader(bytesData)) //用url和bytesData创建一个请求结构体
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

func TestPay(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/payments"
	//配置请求内容
	data := make(map[string]interface{})                                  //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["order_id"] = 12                                               //配置data的内容
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
func TestCreateOrder(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/users/createOrder"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["carts"] = []int{5, 6, 7}
	data["username"] = "小明"
	data["user_id"] = 1
	data["address"] = "榕园"
	data["phone"] = "12345678713"                                         //配置data的内容
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

func TestGetCart(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/users/getCart"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部

	data["user_id"] = 1                                                   //配置data的内容
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
func TestGetOrder(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/users/getorder"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部

	data["user_id"] = 1                                                   //配置data的内容
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
func TestCreateComment(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()

	url := "http://127.0.0.1:3000/apifront/users/comment"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["username"] = "小明"
	data["product_id"] = 1
	data["product_comment"] = "好看"
	data["score"] = "10"                                                  //配置data的内容
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
func TestAddCollect(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/collect/addCollect"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["user_id"] = 1
	data["product_id"] = 1
	                                                 //配置data的内容
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
func TestGetCollect(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/collect/getCollect"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["user_id"] = 1
	
	                                               
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
func TestAddShoppingCart(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/shoppingCart/addShoppingCart"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["user_id"] = "小明"
	data["product_id"] = 1
	                                               
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
func TestUpdateShoppingCart(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/shoppingCart/updateShoppingCart"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["user_id"] = 1
	data["product_id"] = 1
	data["num"]  =1                                       
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
func TestDeleteShoppingCart(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/shoppingCart/deleteShoppingCart"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["user_id"] = "小明"
	data["product_id"] = 1
	                                       
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
func TestGetUserAddress(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/getUseraddress"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["user_id"] = 1
	
	                                       
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
func TestPostUserAddress(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/postUseraddress"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	
	                                       
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
func TestDeleteAddress(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/deleteAddress"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	data["user_id"] = 1
	data["addressID"] = 1
	                                       
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
func TestTestEdit(t *testing.T) {
	//启动一个专用于测试的服务器
	go server()
	time.Sleep(time.Second * 5)
	url := "http://127.0.0.1:3000/apifront/users/saveEdit"
	//配置请求内容
	data := make(map[string]interface{}) //创建一个map类型，从字符串映射到任意类型，作为请求的头部
	
	                                       
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