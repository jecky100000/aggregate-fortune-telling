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

			v := strconv.FormatFloat(amount*config.Rate, 'g', -1, 64)
			order.Des = "八字测算" + v + "元"
			if err := ay.Db.Save(&order).Error; err != nil {
				ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
				return
			}
			if getForm.Type == 1 {
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

		} else {
			ay.Json{}.Msg(c, 400, "支付类型错误", gin.H{})
			return
		}

	}

	if user.Amount < amount {
		ay.Json{}.Msg(c, 406, "余额不足，请选择其他支付方式", gin.H{})
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
