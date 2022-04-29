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

type EnvelopesController struct {
}

type EnvelopesSendForm struct {
	Amount float64 `form:"amount" binding:"required" label:"金额"`
	To     string  `form:"to" binding:"required" label:"对象"`
	Remark string  `form:"remark"`
}

func (con EnvelopesController) Send(c *gin.Context) {
	var getForm EnvelopesSendForm
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

	if user.Amount < getForm.Amount {
		ay.Json{}.Msg(c, 400, "账户余额不足", gin.H{})
		return
	}

	user.Amount = user.Amount - getForm.Amount
	if err := ay.Db.Save(&user).Error; err != nil {
		ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
		return
	}

	var master models.User
	ay.Db.Where("phone = ?", getForm.To).First(&master)

	if master.Id == 0 {
		ay.Json{}.Msg(c, 400, "发送对象不存在", gin.H{})
		return
	}

	oid := ay.MakeOrder(time.Now())

	order := &models.Order{
		Oid:        oid,
		Type:       9,
		Ip:         c.ClientIP(),
		Des:        "发送红包" + strconv.FormatFloat(getForm.Amount, 'g', -1, 64) + "元",
		Amount:     getForm.Amount,
		Uid:        user.Id,
		Status:     0,
		Appid:      Appid,
		OutTradeNo: oid,
		Op:         2,
		OldAmount:  getForm.Amount,
		ToUid:      master.Id,
		Remark:     getForm.Remark,
	}

	ay.Db.Create(order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "数据错误", gin.H{})
	} else {
		ay.Json{}.Msg(c, 200, "发送成功", gin.H{
			"oid":    oid,
			"remark": getForm.Remark,
		})
	}
}

type EnvelopesDetailForm struct {
	Oid float64 `form:"oid" binding:"required" label:"订单号"`
}

func (con EnvelopesController) Detail(c *gin.Context) {
	var getForm EnvelopesDetailForm
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
	ay.Db.Where("oid = ?", getForm.Oid).First(&order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}
	ay.Json{}.Msg(c, 200, "发送成功", gin.H{
		"avatar":   ay.Yaml.GetString("domain") + user.Avatar,
		"remark":   order.Remark,
		"amount":   order.Amount,
		"status":   order.Status,
		"nickname": user.NickName,
	})

}
