/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package api

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type ForceController struct {
}

type forceSendForm struct {
	Amount float64 `form:"amount" binding:"required" label:"金额"`
	To     string  `form:"to" binding:"required" label:"对象"`
	Remark string  `form:"remark"`
}

// Send 发起索要红包
func (con ForceController) Send(c *gin.Context) {
	var getForm forceSendForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	isMaster, msg, master := AuthMaster()
	if !isMaster {
		ay.Json{}.Msg(c, 401, msg, gin.H{})
		return
	}

	var user models.User
	ay.Db.Where("type = 0 and phone = ?", getForm.To).First(&user)
	if user.Id == 0 {
		ay.Json{}.Msg(c, 400, "发送对象不存在", gin.H{})
		return
	}

	tx := ay.Db.Begin()

	oid := ay.MakeOrder(time.Now())

	order := &models.Order{
		Oid:        oid,
		Type:       6,
		Ip:         c.ClientIP(),
		Des:        "大师发送索要红包" + strconv.FormatFloat(getForm.Amount, 'g', -1, 64) + "元",
		Amount:     getForm.Amount,
		Uid:        master.Id,
		Status:     0,
		Appid:      Appid,
		OutTradeNo: oid,
		Op:         2,
		OldAmount:  getForm.Amount,
		ToUid:      user.Id,
		Remark:     getForm.Remark,
	}

	if err := tx.Create(&order).Error; err == nil {
		tx.Commit()
		// 上级消费
		//models.UserInviteConsumptionModel{}.Set(user.Id, user.Pid, getForm.Amount, oid)
		ay.Json{}.Msg(c, 200, "发送成功", gin.H{
			"oid":    oid,
			"remark": getForm.Remark,
		})
	} else {
		tx.Rollback()
		ay.Json{}.Msg(c, 400, "数据错误", gin.H{})
	}
}

type forceDoForm struct {
	Oid       float64 `form:"oid" binding:"required" label:"订单号"`
	PayType   int     `form:"pay_type" binding:"required" label:"支付方式"`
	ReturnUrl string  `form:"return_url" binding:"required" label:"同步地址"`
}

// Do 支付
func (con ForceController) Do(c *gin.Context) {
	var getForm forceDoForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var requestUser models.User
	ay.Db.First(&requestUser, "id = ?", GetToken(Token))

	if requestUser.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var order models.Order
	ay.Db.Where("oid = ?", getForm.Oid).First(&order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}

	if order.Status == 1 {
		ay.Json{}.Msg(c, 400, "该红包已支付过", gin.H{})
		return
	}

	if order.ToUid != requestUser.Id {
		ay.Json{}.Msg(c, 400, "不是您的索要红包", gin.H{})
		return
	}

	var master models.User
	ay.Db.Where("id = ?", order.Uid).First(&master)

	if getForm.PayType == 3 {
		if requestUser.Amount < order.Amount {
			ay.Json{}.Msg(c, 406, "余额不足", gin.H{})
			return
		}

		tx := ay.Db.Begin()

		order.Status = 1

		if err := tx.Save(&order).Error; err == nil {
			master.Amount += order.Amount
			// 增加大师余额
			if err := tx.Save(&master).Error; err == nil {
				// 扣钱用户
				requestUser.Amount -= order.Amount
				if err := tx.Save(&requestUser).Error; err == nil {
					tx.Commit()
					ay.Json{}.Msg(c, 200, "成功支付", gin.H{})
					return
				} else {
					tx.Rollback()
					ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
					return
				}
			} else {
				tx.Rollback()
				ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
				return
			}
		} else {
			tx.Rollback()
			ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
			return
		}
	} else {
		order.ReturnUrl = getForm.ReturnUrl
		ay.Db.Save(&order)

		if Appid == 1 {
			if getForm.PayType == 1 {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"url": ay.Yaml.GetString("domain") + "/pay/alipay?oid=" + order.Oid,
				})
				return
			} else {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"url": ay.Yaml.GetString("domain") + "/pay/wechat?oid=" + order.Oid,
				})
				return
			}
		} else if Appid == 2 {
			is, msg, rj := BaiDuController{}.Baidu(order.Oid)

			if is {
				ay.Json{}.Msg(c, 200, "success", gin.H{
					"info": rj,
				})
				return
			} else {
				ay.Json{}.Msg(c, 400, msg, gin.H{})
				return
			}
		}

	}

}

type forceDetailForm struct {
	Oid float64 `form:"oid" binding:"required" label:"订单号"`
}

// Detail 红包详情
func (con ForceController) Detail(c *gin.Context) {
	var getForm forceDetailForm
	if err := c.ShouldBind(&getForm); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	var requestUser models.User
	ay.Db.First(&requestUser, "id = ?", GetToken(Token))

	if requestUser.Id == 0 {
		ay.Json{}.Msg(c, 401, "Token错误", gin.H{})
		return
	}

	var order models.Order
	ay.Db.Where("oid = ?", getForm.Oid).First(&order)

	if order.Id == 0 {
		ay.Json{}.Msg(c, 400, "订单不存在", gin.H{})
		return
	}

	var master models.User
	ay.Db.Where("id = ?", order.Uid).First(&master)

	var user models.User
	ay.Db.Where("id = ?", order.ToUid).First(&user)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"master_avatar":   ay.Yaml.GetString("domain") + master.Avatar,
		"master_nickname": master.NickName,
		"master_phone":    master.Phone,
		"user_avatar":     ay.Yaml.GetString("domain") + user.Avatar,
		"user_nickname":   user.NickName,
		"remark":          order.Remark,
		"amount":          order.Amount,
		"status":          order.Status,
	})
}
