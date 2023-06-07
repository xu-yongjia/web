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
	kPrivateKey = "MIIEpAIBAAKCAQEAyZMA4UNDthsHLOEG5lAYtioN+RTM7vAbHV9Ud9UnI6vhEvmXQzNYN4Vz9jKbsDzFZoZSAbe2e65zyVkn/vLFXQY2AnstoMnPu0rzLT3gCuREOeTAtsFF3riBf40jTubov0UukSp1+mehQuhsicH8bu+QyoFTDxhA7KdSm3eCG0/RhQiHPFtCJYGJKMjf9/NWpYkdqF//ugS3zzVnftB5tDdEXWmUbYXRocnfjqtUGDuXknTXOWXO2hhihU8feyLIwmwxlz3CCCBn1xjIAJXRGk7sdxEhUgWQ3a3cLPXnwchKaVpfLOgl1Lc+Sj9mhLpCB5+IUc9uxL4GyKxkd89itQIDAQABAoIBABlZpXiu7usOBx77xdRdKl5ud4dluEgsZhyRipvvcJap3Qp33TAUWvnQjL7fqnjFi8XTCqkwo4B2dhVSYxdLDMsgJta1Cw33k2poiKCd8XktL+9f92SEdfcGmTKjW1dm7spvNuYGwwOU2NEDTfMdNFP3SOxOy1VQpYmdmWP36aLpc7otjlsY+tw6ys4jby/sVAtxCGOT0edmxTFGuTSNwKRtjSoyZGe/O83HHWuyf0txkT6FcC9Yb5jPnTnAgLbmR59xWVuIkm4VUJaWvJyd1Gtzr5GnycHJWe4OfiBCoahKERbf94MMGv95V305Rq6XzmVmU9rJFpG7OcwUSJCNX8ECgYEA5Xxdda/h2FGdzykZx5N1IFtNsm+DN+PPP5Q9o068GSW/I8PbWCzTNAEV+8CX4bAhKFJxPfNQ4h8cNyH0Tm/t2iFqtSvTpCbttnz68l1U8x5SAfBkKAyKotzMO9QMZxtZa1OVfJQqbkWmozKB+QJsIobWmtY28sLP41ymsuZx3ikCgYEA4N0U5+hPPBG16U5STsSrz0z+e8oYLyr7nd/G7Kajd2lQoZINgJTSEBXFnoPYFzcQKcvTrn6nkBHK2Cp14wPoYg7eDWHZn/C8dLmeA8Bf1U9J6gywuB6ur2duA4fqAIdByn5o7yz5Kk8hyIwxWSMu2HYjxYEoQIhN0bFI04WgWa0CgYEAnmVX0OpkVT4IUbvJunIyLvI2T8oeZQ5Nt1GntbebpzPhmVsPY+4UnVl+hfTqIHHEgSNeXGc8VccpX3iFOzqQjrXwBDrK6wufmgzr3SvEfcZYzlPRnb8CVeVW5pi3MteywZzAmmteWhyjJxOumGQupSoyqE7hQHBssz75JJRg1fkCgYEArrw91aRZhhVBJgY9wYXk+Ncc3jcCKg0dr+XH3/fyCQXmqVckdOmwft1cTrynfArG3aPNX8h+D180IFyGdR2BREnHVeDWywm4eIARAVtJXSobBuxLmxgEqLsX61h/h1+TyXYHtuxcxvrtBrmiU2EaswzC6WRfUh5ZUnyQWgZ5v5UCgYBt8klPnF/iAtSFp4bNWeyKwBv651GqknQFghKDNaX5KIzrKxupwGAw8WWScXglI37WdsGnQ0IUQleZzHHYDmDzyB2l5V4J4+LmIvG0UZErBdIrgqwluTD9g93LxJuQbE/xasXXJBO6v6mpny7IZyamNP0gAkHyZ3clYcbscTs8Og=="
	kServerPort = "9989"
	//设置回调地址域名，将内网穿透的URL放到这里
	kServerDomain = "https://43afa151.r10.cpolar.top"
)

func InitPay() {
	var err error

	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}
	err = client.SetEncryptKey("VP1UE6HugWdMaxH4IWVvVA==")
	if err != nil {
		log.Println(err)
	}
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

	if err = client.SetEncryptKey("FtVd5SgrsUzYQRAPBmejHQ=="); err != nil {
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
		err := client.SetEncryptKey("VP1UE6HugWdMaxH4IWVvVA==") //在支付宝-控制台-沙箱-加密-开发信息-接口内容加密模式，复制后粘贴到这里
		if err != nil {
			log.Println(err)
		}
		url, _ := client.TradePagePay(p)
		fmt.Print(url.String())
		c.Redirect(http.StatusTemporaryRedirect, url.String())
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
