package api1

import (
	"fmt"
	"gintest/DBstruct"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
)

//支付沙箱账号
//juhhys5374@sandbox.com
//111111

var client *alipay.Client

const (
	//应用号（APPID)
	kAppId = "2021000122694478"
	//应用私钥，在支付宝-控制台-沙箱-加密-开发信息-接口加签模式-证书模式-查看，复制后黏贴到这里
	kPrivateKey = "MIIEowIBAAKCAQEAgkfE8a2B8iKEICE9l3KQQSVY6Q6WPFLtGKDBWkpFeWUZpoCsX9B4e92XxL8leCgjGw6KJ4kvbnd88/8wexqj8AF1TruCR1xd1ihvBhPkCOF18aVNBarSsLzW3qvhpwBZvPKxm+m6WgYOt8cM6pQ7VUBtbE016GIbq8qgsjPb3mxMFgL9/uq5kHV3/bWExhw4jmPSZR/6/4elB83Vh6NsN8Pu3t5cMJDByrngNAR6XOog87CS9abjQrGv4vFxvRR+S+v5vT1nCdtSSCN31YDNa30NVrnaKXng+9h8pIL2m3D6Sx+qqbL+ZL5r5cMVwT35RS0q5yiFpbUgupF76XBOAQIDAQABAoIBACylYbWcy0pqT7MfERdkeFEK9nQvvoZ/uKbW+Evz5ekbH3Yy5F+VdqSkJM2U38dUoes0yV2RIsMNibGoQzn8wC+QoF3R3myMF9A7XQjNKjls4llkV4fYqYlNvajjaFotwZMB7KHMsEMHjmWmfiOplse5Y4ZY01+Gt4rgzr8jDHclT968R2nSEdscKABnfeTU5Hzh5NBNAvsto70Ph05zIU7yR/pUr5sgE9qB3YloR8qZfnVbqTiwo/q012a4YWG4+OWiLRGgymUXtT42ikQgHMA7rZtKqASh7eZ446hNedtIQGY8nQGhRVOcwLp8QUc9eFpBe8uZVHq6Klq/gzSZiRECgYEA4WDZdWwQ5RzPPwytDMcQD7PT2PZgGi9evzabvVrOrLDE+WFC5mtuZQWtUGk+5KvVMBHqzGZ9xB1H6amze+4ejXsuxqDDoqF6A3bD4LolIikNBfDaA/kQezTGftfqZoVr6qFFE4MdVFb9WbYbaBXqVSSz32DTQ4vrVZ6KI/YRte0CgYEAk/s0Yho27O8MwzMVAPUiGSBXB9l8zk+lO+nwWtWGS3GSQ2SLslJAekMzM8emlBMSZf4HOC/2JrAitACX/wRtYQb2DC/67IVDSqtxBkrIex6Yft3Re3zRGHH+cQCMjAT9NhIvRHlqKvvqH+4aJTM+8i0Dmrt4vbHNqBUt/x7YteUCgYEAp8B6m5fQC3iiGX/J3lTcc5OZ/0eCyHqxvwlaalqIsqjwM95c+TzFcpQDKHLtyS7NA3aSmYk/2xUgTZ9uZLQcMZqaemV2uDoQoztnGIH02bRMaTuYSpS727iwdgEoCx8L8BwRWjChSAFnmbqeM578SFuFNqLPHNusUpIqqBGjpRECgYBSaf3jgPwD+qJd+A4DiYIjAjrWGU8Dy6GXe5Mr021GiQrPU6jleB+Gt/RFKzTflLuTZ+V3amapZ0D8hLQiB8Iu3mSbpwEVRqaSaGa6/V0oDIMbbMglAP+moJ33KzobJRAsU/ZVKVVyePgMLlwLSb4Nu/oc67mz09lYeeszAWfp7QKBgGycHccJ2x2/oNB1nZqjKbgrVvtE7cJIlug6pxSdp0oVhLWU0vSkjcEIciTRiarWQ7OESb47y25eErEv4v9Mo2iRAalYwPn6QJPW/bvo/u2HnbF+Hi+0b8etqstQ9WQGUNwzRmZ3A90z9GLJxf551WIqx3dAnR9cpXhzBgXzxSTC"
	kServerPort = "9989"
	//设置回调地址域名，将内网穿透的URL放到这里
	kServerDomain = "http://396b01f1.r3.cpolar.top"
)

func InitPay() {
	var err error

	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}
	// err = client.SetEncryptKey("yhL/DROt5T1GJubzpHy2eA==")
	// if err != nil {
	// 	log.Println(err)
	// }
	// 加载证书
	//在支付宝-控制台-沙箱-加密-开发信息-接口加签模式-证书模式-查看，可以下载三份证书，分别放到cert文件夹中
	if err = client.LoadAppCertPublicKeyFromFile("cert/appPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAliPayRootCertFromFile("cert/alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAlipayCertPublicKeyFromFile("cert/alipayPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}

	if err = client.SetEncryptKey("yhL/DROt5T1GJubzpHy2eA=="); err != nil {
		log.Println("加载内容加密密钥发生错误", err)
		return
	}
}

func Pay(c *gin.Context) {
	type msg struct {
		Order_id uint64 `json:"order_id"`
	}
	var m msg
	price := 0.0
	if e := c.ShouldBindJSON(&m); e == nil {
		Order_list := []DBstruct.Order{}
		if e = DBstruct.DB.Where("order_id = ? AND status = ?", m.Order_id, "未支付").Find(&Order_list).Error; e == nil && len(Order_list) > 0 {
			fmt.Println(Order_list)
			for _, order := range Order_list {
				product := DBstruct.Product{}
				if e = DBstruct.DB.Where("id = ?", order.ProductId).Find(&product).Error; e == nil {
					productprice, _ := strconv.ParseFloat(product.DiscountPrice, 64)
					fmt.Println(float64(order.Num), productprice)
					price += float64(order.Num) * productprice
				} else {
					c.JSON(201, ERRRESPONSE("找不到商品", 201))
					return
				}
			}
		} else {
			c.JSON(201, ERRRESPONSE("不存在具有该订单号的未支付订单", 201))
			return
		}
		var tradeNo = fmt.Sprintf("%d", m.Order_id)
		var p = alipay.TradePagePay{}
		p.NotifyURL = kServerDomain + "/api/v1/alipay/notify"   //回调通知路径
		p.ReturnURL = kServerDomain + "/api/v1/alipay/callback" //回调返回路径
		p.Subject = "支付测试:" + tradeNo
		p.OutTradeNo = tradeNo
		fmt.Printf("%.2f", price)
		p.TotalAmount = fmt.Sprintf("%.2f", price) //价格（两位小数，不能为0）
		p.ProductCode = "FAST_INSTANT_TRADE_PAY"
		// err := client.SetEncryptKey("yhL/DROt5T1GJubzpHy2eA==") //在支付宝-控制台-沙箱-加密-开发信息-接口内容加密模式，复制后粘贴到这里
		// if err != nil {
		// 	log.Println(err)
		// }
		url, _ := client.TradePagePay(p)
		fmt.Print(url.String())
		c.JSON(200, SUCCESSRESPONSE(url.String()))
	} else { //解析json失败
		c.JSON(400, ERRRESPONSE(e.Error(), 201))
	}
}

func Callback(c *gin.Context) {
	c.Request.ParseForm()

	err := client.VerifySign(c.Request.Form)
	if err != nil {
		log.Println("回调验证签名发生错误", err)
		c.String(http.StatusBadRequest, "回调验证签名发生错误")
		return
	}

	log.Println("回调验证签名通过")

	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	log.Printf("订单 %s 支付成功 \n", outTradeNo)
	paiedOrderList := []DBstruct.Order{}
	DBstruct.DB.Where("order_id=?", outTradeNo).Find(&paiedOrderList)
	for _, paiedOrder := range paiedOrderList {
		product := DBstruct.Product{}
		product.ID = uint(paiedOrder.ProductId)
		DBstruct.DB.Where("id=?", paiedOrder.ProductId).First(&product)
		product.Sales = product.Sales + paiedOrder.Num
		DBstruct.DB.Save(&product)
	}
	DBstruct.DB.Model(DBstruct.Order{}).Where("order_id = ?", outTradeNo).Update("status", "已支付")
	rsp, err := client.TradeQuery(p)
	if err != nil {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		return
	}

	if rsp.IsFailure() {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "https://baidu.com")
	//c.String(http.StatusOK, "订单 %s 支付成功", outTradeNo)
}

func Notify(c *gin.Context) {
	var noti, err = client.DecodeNotification(c.Request)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		return
	}

	log.Println("解析异步通知成功:", noti.NotifyId)

	var p = alipay.NewPayload("alipay.trade.query")
	p.AddField("out_trade_no", noti.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = client.Request(p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", noti.OutTradeNo, err.Error())
		return
	}
	if rsp.IsFailure() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", noti.OutTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}
	if rsp.TradeStatus == "TRADE_SUCCESS" {
		log.Printf("订单 %s 支付成功 \n", noti.OutTradeNo)
		paiedOrderList := []DBstruct.Order{}
		DBstruct.DB.Where("order_id=?", noti.OutTradeNo).Find(&paiedOrderList)
		for _, paiedOrder := range paiedOrderList {
			product := DBstruct.Product{}
			product.ID = uint(paiedOrder.ProductId)
			DBstruct.DB.Where("id=?", paiedOrder.ProductId).First(&product)
			product.Sales = product.Sales + paiedOrder.Num
			DBstruct.DB.Save(&product)
		}
		DBstruct.DB.Model(DBstruct.Order{}).Where("order_id = ?", noti.OutTradeNo).Update("status", "已支付")
		//这里需要通知商家，或者由商家主动查询
	}
	client.ACKNotification(c.Writer)
}
