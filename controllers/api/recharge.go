/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
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

	if Appid == 1 {
		// web支付
		res, order := con.MakeOrder(user.Id, getForm.Amount, getForm.ReturnUrl, GetRequestIP(c), getForm.Type)

		if res == 0 {
			ay.Json{}.Msg(c, 400, "订单创建失败", gin.H{})
			return
		}
		if getForm.Type == 1 {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"url": ay.Yaml.GetString("domain") + "/pay/alipay?oid=" + order,
			})
		} else {
			ay.Json{}.Msg(c, 200, "success", gin.H{
				"url": ay.Yaml.GetString("domain") + "/pay/wechat?oid=" + order,
			})
		}

	}
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
