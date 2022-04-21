/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"context"
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/wechat"
	"strconv"
	"time"
)

type RechargeController struct {
}

type GetRechargeMainForm struct {
	Type      int     `form:"type" binding:"required" label:"类型"`
	Amount    float64 `form:"amount" binding:"required" label:"金额"`
	ReturnUrl string  `form:"return_url" binding:"required" label:"同步地址"`
}

func (con RechargeController) Main(c *gin.Context) {
	var getForm GetRechargeMainForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if getForm.Amount < 1 {
		ay.Json{}.Msg(c, 400, "充值不能小于1元", gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var (
		code int
		msg  string
	)

	if Appid == 1 {
		// web支付
		res, order := con.MakeOrder(user.Id, getForm.Amount, getForm.ReturnUrl, GetRequestIP(c), getForm.Type)

		code, msg = con.Web(order, getForm.Type, getForm.Amount, getForm.ReturnUrl, GetRequestIP(c))

		if res == 0 {
			ay.Json{}.Msg(c, 400, "订单创建失败", gin.H{})
			return
		}
		if code == 1 {
			if getForm.Type == 1 {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"url": ay.Yaml.GetString("domain") + "/pay/alipay?oid=" + order,
				})
			} else {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"url": ay.Yaml.GetString("domain") + "/pay/wechat?oid=" + order,
				})
			}

		} else {
			ay.Json{}.Msg(c, 400, msg, gin.H{})
		}
	}
}

func (con RechargeController) Web(oid string, payType int, amount float64, returnUrl string, ip string) (int, string) {

	config := models.ConfigModel{}.GetId(1)

	ctx, _ := context.WithCancel(context.Background())

	if payType == 1 {

		var pay models.Pay
		ay.Db.First(&pay, "id = ?", 7)

		client, err := alipay.NewClient(pay.Appid, pay.VKey, true)
		if err != nil {
			return 0, err.Error()
		}
		client.SetLocation(alipay.LocationShanghai).
			SetCharset(alipay.UTF8).                                         // 设置字符编码，不设置默认 utf-8
			SetSignType(alipay.RSA2).                                        // 设置签名类型，不设置默认 RSA2
			SetReturnUrl(returnUrl).                                         // 设置返回URL
			SetNotifyUrl(ay.Yaml.GetString("domain") + "/api/notify/alipay") // 设置异步通知URL

		bm := make(gopay.BodyMap)
		v := strconv.FormatFloat(amount*config.Rate, 'g', -1, 64)
		bm.Set("subject", "充值"+v+"元").
			Set("product_code", "QUICK_WAP_PAY").
			Set("out_trade_no", oid).
			Set("total_amount", amount).
			Set("quit_url", returnUrl) // 中途退出

		aliRsp, err := client.TradeWapPay(ctx, bm)

		if err != nil {
			return 0, err.Error()
		}
		var order models.Order
		ay.Db.First(&order, "oid = ?", oid)
		order.Json = aliRsp
		ay.Db.Save(&order)

		return 1, ""
	} else if payType == 2 {
		// 微信支付 jsapi需要跳转页面获取openid
		var pay models.Pay
		ay.Db.First(&pay, "id = ?", 6)

		client := wechat.NewClient(pay.Appid, pay.MchId, pay.VKey, true)
		// 打开Debug开关，输出请求日志，默认关闭
		//client.DebugSwitch = gopay.DebugOn
		client.SetCountry(wechat.China)

		bm := make(gopay.BodyMap)
		v := strconv.FormatFloat(amount*config.Rate, 'g', -1, 64)
		bm.Set("nonce_str", util.GetRandomString(32)).
			Set("body", "充值"+v+"元").
			Set("out_trade_no", oid).
			Set("total_fee", amount*100).
			Set("spbill_create_ip", ip).
			Set("notify_url", ay.Yaml.GetString("domain")+"/api/notify/wechat").
			Set("trade_type", "MWEB").
			Set("device_info", "WEB").
			Set("sign_type", "MD5").
			SetBodyMap("scene_info", func(bm gopay.BodyMap) {
				bm.SetBodyMap("h5_info", func(bm gopay.BodyMap) {
					bm.Set("type", "Wap")
					bm.Set("wap_url", ay.Yaml.GetString("domain"))
					bm.Set("wap_name", "H5测试支付")
				})
			}) /*.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")*/

		wxRsp, err := client.UnifiedOrder(ctx, bm)

		if err != nil {
			return 0, err.Error()
		}

		var order models.Order
		ay.Db.First(&order, "oid = ?", oid)
		order.Json = wxRsp.MwebUrl
		ay.Db.Save(&order)
		return 1, ""
	} else {
		return 0, "支付类型不正确"
	}

	return 0, ""
}

// MakeOrder 生成订单号
func (con RechargeController) MakeOrder(uid int64, amount float64, return_url string, ip string, PayType int) (int, string) {
	oid := ay.MakeOrder(time.Now())

	v := strconv.FormatFloat(amount, 'g', -1, 64)

	order := &models.Order{
		Oid:        oid,
		Type:       9,
		Ip:         ip,
		Des:        "充值" + v + "元",
		Amount:     amount,
		Uid:        uid,
		Status:     0,
		Appid:      Appid,
		PayType:    PayType,
		OutTradeNo: oid,
		Op:         1,
		OldAmount:  amount,
		ReturnUrl:  return_url,
	}

	ay.Db.Create(order)

	if order.Id == 0 {
		return 0, "数据错误，请联系管理员"
	} else {
		return 1, oid
	}
}
