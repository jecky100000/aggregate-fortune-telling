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
	"log"
	"strconv"
	"strings"
	"time"
)

type PayController struct {
}

type GetPayDoForm struct {
	AmountId  int64  `form:"amount_id"`
	Oid       string `form:"oid" binding:"required" label:"订单号"`
	Coupon    int64  `form:"coupon"`
	Return    int    `form:"return"`
	Type      int    `form:"type" binding:"required" label:"类型"`
	ReturnUrl string `form:"return_url" binding:"required" label:"返回地址"`
}

// Do 统一支付
func (con PayController) Do(c *gin.Context) {
	var getForm GetPayDoForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var user models.User
	ay.Db.First(&user, "id = ?", GetToken(Token))

	if user.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var order models.Order
	ay.Db.First(&order, "oid = ?", getForm.Oid)
	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}

	if order.Status == 1 {
		ay.Json{}.Msg(c, 400, "该笔订单已支付过", gin.H{})
		return
	}

	order.ReturnUrl = getForm.ReturnUrl

	config := models.ConfigModel{}.GetId(1)

	// 获取金额
	VAmount := 0.00
	if getForm.Return == 0 && order.Type == 1 {
		var haulAmount models.HaulAmount
		ay.Db.First(&haulAmount, "id = ?", getForm.AmountId)

		if haulAmount.Id == 0 {
			ay.Json{}.Msg(c, 400, "金额错误", gin.H{})
			return
		}
		VAmount = haulAmount.Amount
	} else if order.Type == 1 && getForm.Return == 1 {
		// 八字
		VAmount = config.HaulAmount
	} else {
		VAmount = order.Amount
	}

	// 订单历史金额
	order.OldAmount = VAmount

	// 优惠卷金额
	couponAmount := 0.00
	if getForm.Coupon != 0 && getForm.Return == 0 {
		// 穷逼优惠卷支付
		var coupon models.Coupon
		ay.Db.First(&coupon, "id = ? and uid = ?", getForm.Coupon, user.Id)

		if coupon.Id == 0 {
			ay.Json{}.Msg(c, 400, "优惠卷不存在", gin.H{})
			return
		}

		couponTypeArr := strings.Split(coupon.Product, ",")

		vType := 0
		for _, v := range couponTypeArr {
			cl, _ := strconv.Atoi(v)
			if cl == order.Type {
				vType = 1
			}
		}
		if vType == 0 {
			ay.Json{}.Msg(c, 400, "优惠卷不适用于此产品", gin.H{})
			return
		}

		if coupon.AmountThan > VAmount {
			ay.Json{}.Msg(c, 400, "优惠卷不适用于此产品，金额错误", gin.H{})
			return
		}

		if coupon.EffectiveAt.Unix() < time.Now().Unix() {
			ay.Json{}.Msg(c, 400, "优惠卷已过期", gin.H{})
			return

		}

		if coupon.Status != 0 {
			ay.Json{}.Msg(c, 400, "优惠卷已使用", gin.H{})
			return

		}

		order.Coupon = coupon.Id
		couponAmount = coupon.Amount
	}

	// 减少用户余额
	amount := 0.00
	if getForm.Return == 0 {
		amount = VAmount - couponAmount
	} else {
		amount = VAmount - order.Discount
	}

	order.Amount = amount

	if getForm.Type != 3 {
		order.PayType = getForm.Type
		if err := ay.Db.Save(&order).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
			return
		}
		if getForm.Type == 1 || getForm.Type == 2 {
			// 支付宝 微信
			code, msg := con.Web(order.OutTradeNo, getForm.Type, order.Amount, order.ReturnUrl, GetRequestIP(c))

			if code == 1 {
				if getForm.Type == 1 {
					ay.Json{}.Msg(c, 200, "success", gin.H{
						"url": ay.Yaml.GetString("domain") + "/pay/alipay?oid=" + order.OutTradeNo,
					})
					return
				} else {
					ay.Json{}.Msg(c, 200, "success", gin.H{
						"url": ay.Yaml.GetString("domain") + "/pay/wechat?oid=" + order.OutTradeNo,
					})
					return
				}

			} else {
				ay.Json{}.Msg(c, 400, msg, gin.H{})
				return
			}

		} else {
			ay.Json{}.Msg(c, 400, "支付类型错误", gin.H{})
			return
		}

		log.Println(order.Amount, order.OldAmount, order.Coupon)
		return
	}

	if user.Amount < amount {
		ay.Json{}.Msg(c, 406, "余额不足", gin.H{})
		return
	}

	user.Amount = user.Amount - amount
	ay.Db.Save(&user)

	models.UserInviteConsumptionModel{}.Set(user.Id, user.Pid, amount, order.Oid)

	// 订单设置已支付
	order.Status = 1
	order.PayType = 9
	order.PayTime = time.Now().Format("2006-01-02 15:04:05")
	if err := ay.Db.Save(&order).Error; err != nil {
		ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
		return
	}

	// 优惠卷设置过期
	if getForm.Coupon != 0 {
		var coupon models.Coupon
		ay.Db.First(&coupon, "id = ?", getForm.Coupon)
		coupon.Status = 1
		coupon.UsedAt = time.Now().Format("2006-01-02 15:04:05")
		if err := ay.Db.Save(&coupon).Error; err != nil {
			ay.Json{}.Msg(c, 400, "修改失败", gin.H{})
			return
		}
	}

	ay.Json{}.Msg(c, 200, "支付成功", gin.H{})

}

func (con PayController) Web(oid string, payType int, amount float64, returnUrl string, ip string) (int, string) {

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
		bm.Set("subject", "八字测算支付"+v+"元").
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
		bm.Set("nonce_str", util.RandomString(32)).
			Set("body", "八字测算支付"+v+"元").
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
