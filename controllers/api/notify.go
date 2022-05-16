/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"encoding/json"
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/wechat"
	"log"
	"net/http"
	"time"
)

type NotifyController struct {
}

func (con NotifyController) AliPay(c *gin.Context) {

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", 7)

	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request)
	con.CheckErr(err)

	j, err := json.Marshal(notifyReq)
	con.CheckErr(err)

	ay.CreateMutiDir("log/alipay")
	ay.WriteFile("log/alipay/"+notifyReq.Get("out_trade_no")+".txt", []byte(j), 0755)

	ok, err := alipay.VerifySign(pay.PayKey, notifyReq)

	if !ok {
		log.Println(err)
		c.String(http.StatusOK, "%s", "fail")
		return
	}

	var order models.Order
	ay.Db.First(&order, "out_trade_no = ?", notifyReq.Get("out_trade_no"))

	// 查询订单失败
	if order.Id == 0 {
		c.String(http.StatusOK, "%s", "fail")
		return
	}

	// 订单已处理过
	if order.Status == 1 {
		c.String(http.StatusOK, "%s", "success")
		return
	}

	res := 0

	switch order.Type {
	case 9:
		// 增加余额
		res = con.AddUserAmount(order.Uid, order.Amount)
	case 1:
		res = 1
	}

	if res == 1 {
		// 优惠卷设置过期
		var coupon models.Coupon
		ay.Db.First(&coupon, "id = ?", order.Coupon)
		if coupon.Id != 0 {
			coupon.Status = 1
			coupon.UsedAt = time.Now().Format("2006-01-02 15:04:05")
			ay.Db.Save(&coupon)
		}
		order.Status = 1
		order.PayType = 1
		order.TradeNo = notifyReq.Get("trade_no")
		order.PayTime = time.Now().Format("2006-01-02 15:04:05")
		ay.Db.Save(&order)
		c.String(http.StatusOK, "%s", "success")
	} else {
		c.String(http.StatusOK, "%s", "fail")
	}

}

func (con NotifyController) WeChat(c *gin.Context) {

	var pay models.Pay
	ay.Db.First(&pay, "id = ?", 6)

	notifyReq, err := wechat.ParseNotifyToBodyMap(c.Request)
	con.CheckErr(err)
	log.Println(notifyReq)

	j, err := json.Marshal(notifyReq)
	con.CheckErr(err)

	ay.CreateMutiDir("log/wechat")
	ay.WriteFile("log/wechat/"+notifyReq.Get("out_trade_no")+".txt", []byte(j), 0755)

	// 验签操作
	ok, err := wechat.VerifySign(pay.VKey, wechat.SignType_MD5, notifyReq)

	if !ok {
		log.Println(err)
		c.String(http.StatusOK, "%s", "fail")
		return
	}

	var order models.Order
	ay.Db.First(&order, "out_trade_no = ?", notifyReq.Get("out_trade_no"))

	// 查询订单失败
	if order.Id == 0 {
		rsp := new(wechat.NotifyResponse) // 回复微信的数据
		rsp.ReturnCode = gopay.FAIL
		rsp.ReturnMsg = gopay.FAIL
		c.String(http.StatusOK, "%s", rsp.ToXmlString())
		return
	}

	// 订单已处理过
	if order.Status == 1 {
		rsp := new(wechat.NotifyResponse) // 回复微信的数据
		rsp.ReturnCode = gopay.SUCCESS
		rsp.ReturnMsg = gopay.OK
		c.String(http.StatusOK, "%s", rsp.ToXmlString())
		return
	}

	res := 0

	switch order.Type {
	case 9:
		// 增加余额
		res = con.AddUserAmount(order.Uid, order.Amount)
	case 1:
		res = 1
	}

	if res == 1 {
		// 优惠卷设置过期
		var coupon models.Coupon
		ay.Db.First(&coupon, "id = ?", order.Coupon)
		if coupon.Id != 0 {
			coupon.Status = 1
			coupon.UsedAt = time.Now().Format("2006-01-02 15:04:05")
			ay.Db.Save(&coupon)
		}

		order.Status = 1
		order.PayType = 2
		order.TradeNo = notifyReq.Get("transaction_id")
		order.PayTime = time.Now().Format("2006-01-02 15:04:05")
		ay.Db.Save(&order)
		rsp := new(wechat.NotifyResponse) // 回复微信的数据
		rsp.ReturnCode = gopay.SUCCESS
		rsp.ReturnMsg = gopay.OK
		c.String(http.StatusOK, "%s", rsp.ToXmlString())
	} else {
		rsp := new(wechat.NotifyResponse) // 回复微信的数据
		rsp.ReturnCode = gopay.FAIL
		rsp.ReturnMsg = gopay.FAIL
		c.String(http.StatusOK, "%s", rsp.ToXmlString())
	}

}

// AddUserAmount 增加用户余额
func (con NotifyController) AddUserAmount(uid int64, amount float64) int {
	// 查询用户
	var user models.User
	ay.Db.First(&user, "id = ?", uid)

	if user.Id == 0 {
		log.Println("用户未找到")
		return 0
	}

	user.Amount = user.Amount + amount
	ay.Db.Save(&user)

	return 1

}

func (con NotifyController) CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
