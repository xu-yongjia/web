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
	kPrivateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDJkwDhQ0O2Gwcs4QbmUBi2Kg35FMzu8BsdX1R31Scjq+ES+ZdDM1g3hXP2MpuwPMVmhlIBt7Z7rnPJWSf+8sVdBjYCey2gyc+7SvMtPeAK5EQ55MC2wUXeuIF/jSNO5ui/RS6RKnX6Z6FC6GyJwfxu75DKgVMPGEDsp1Kbd4IbT9GFCIc8W0IlgYkoyN/381aliR2oX/+6BLfPNWd+0Hm0N0RdaZRthdGhyd+Oq1QYO5eSdNc5Zc7aGGKFTx97IsjCbDGXPcIIIGfXGMgAldEaTux3ESFSBZDdrdws9efByEppWl8s6CXUtz5KP2aEukIHn4hRz27EvgbIrGR3z2K1AgMBAAECggEAGVmleK7u6w4HHvvF1F0qXm53h2W4SCxmHJGKm+9wlqndCnfdMBRa+dCMvt+qeMWLxdMKqTCjgHZ2FVJjF0sMyyAm1rULDfeTamiIoJ3xeS0v71/3ZIR19waZMqNbV2buym825gbDA5TY0QNN8x00U/dI7E7LVVCliZ2ZY/fpoulzui2OWxj63DrKziNvL+xUC3EIY5PR52bFMUa5NI3ApG2NKjJkZ787zccda7J/S3GRPoVwL1hvmM+dOcCAtuZHn3FZW4iSbhVQlpa8nJ3Ua3OvkafJwclZ7g5+IEKhqEoRFt/3gwwa/3lXfTlGrpfOZWZT2skWkbs5zBRIkI1fwQKBgQDlfF11r+HYUZ3PKRnHk3UgW02yb4M3488/lD2jTrwZJb8jw9tYLNM0ARX7wJfhsCEoUnE981DiHxw3IfROb+3aIWq1K9OkJu22fPryXVTzHlIB8GQoDIqi3Mw71AxnG1lrU5V8lCpuRaajMoH5Amwihtaa1jbyws/jXKay5nHeKQKBgQDg3RTn6E88EbXpTlJOxKvPTP57yhgvKvud38bspqN3aVChkg2AlNIQFcWeg9gXNxApy9OufqeQEcrYKnXjA+hiDt4NYdmf8Lx0uZ4DwF/VT0nqDLC4Hq6vZ24Dh+oAh0HKfmjvLPkqTyHIjDFZIy7YdiPFgShAiE3RsUjThaBZrQKBgQCeZVfQ6mRVPghRu8m6cjIu8jZPyh5lDk23Uae1t5unM+GZWw9j7hSdWX6F9OogccSBI15cZzxVxylfeIU7OpCOtfAEOsrrC5+aDOvdK8R9xljOU9GdvwJV5VbmmLcy17LBnMCaa15aHKMnE66YZC6lKjKoTuFAcGyzPvkklGDV+QKBgQCuvD3VpFmGFUEmBj3BheT41xzeNwIqDR2v5cff9/IJBeapVyR06bB+3VxOvKd8Csbdo81fyH4PXzQgXIZ1HYFEScdV4NbLCbh4gBEBW0ldKhsG7EubGASouxfrWH+HX5PJdge27FzG+u0GuaJTYRqzDMLpZF9SHllSfJBaBnm/lQKBgG3ySU+cX+IC1IWnhs1Z7IrAG/rnUaqSdAWCEoM1pfkojOsrG6nAYDDxZZJxeCUjftZ2wadDQhRCV5nMcdgOYPPIHaXlXgnj4uYi8bRRkSsF0iuCrCW5MP2D3cvEm5BsT/FqxdckE7q/qamfLshnJqY0/SACQfJndyVhxuxxOzw6"
	kServerPort = "9989"
	//设置回调地址域名，将内网穿透的URL放到这里
	kServerDomain = "http://2294e99e.r11.cpolar.top"
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
		if err = DBstruct.DB.Model(DBstruct.Order{}).Where("order_id = ?", noti.OutTradeNo).Update("status", "已支付").Error; err != nil {
			log.Println(err)
		}
		//这里需要通知商家，或者由商家主动查询
	}
	client.ACKNotification(c.Writer)
}
