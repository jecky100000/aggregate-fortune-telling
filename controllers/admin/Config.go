/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package admin

import (
	"gin/ay"
	"gin/models"
	"github.com/gin-gonic/gin"
	"time"
)

type ConfigController struct {
}

// Detail 详情
func (con ConfigController) Detail(c *gin.Context) {

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.Config

	ay.Db.First(&user, 1)

	ay.Json{}.Msg(c, 200, "success", gin.H{
		"info": user,
	})
}

type configOption struct {
	HaulAmount        float64 `form:"haul_amount"`
	Coupon            int     `form:"coupon"`
	CouponName        string  `form:"coupon_name"`
	CouponSubName     string  `form:"coupon_sub_name"`
	CouponProhibit    string  `form:"coupon_prohibit"`
	CouponEffective   string  `form:"coupon_effective"`
	CouponEffectiveAt string  `form:"coupon_effective_at"`
	CouponAmount      float64 `form:"coupon_amount"`
	CouponProduct     string  `form:"coupon_product"`
	CouponAmountThan  float64 `form:"coupon_amount_than"`
	Kf                string  `form:"kf"`
	Keywords          string  `form:"keywords"`
	Description       string  `form:"description"`
	Notice            string  `form:"notice"`
	Praise            string  `form:"praise"`
	Rate              float64 `form:"rate"`
	MasterLink        string  `form:"master_link"`
	UserRegCoupon     int     `form:"user_reg_coupon"`
	HaulDiscount      string  `form:"haul_discount"`
	InviteRate        float64 `form:"invite_rate"`
	WithdrawAmount    float64 `form:"withdraw_amount"`
}

// Option 添加 编辑
func (con ConfigController) Option(c *gin.Context) {
	var data configOption
	if err := c.ShouldBind(&data); err != nil {
		ay.Json{}.Msg(c, 400, ay.Validator{}.Translate(err), gin.H{})
		return
	}

	if Auth() == false {
		ay.Json{}.Msg(c, 401, "请登入", gin.H{})
		return
	}

	var user models.Config
	ay.Db.First(&user, 1)

	user.HaulAmount = data.HaulAmount
	user.Coupon = data.Coupon
	user.CouponName = data.CouponName
	user.CouponSubName = data.CouponSubName
	user.CouponProhibit = data.CouponProhibit
	user.CouponEffective = data.CouponEffective
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", data.CouponEffectiveAt, time.Local)
	user.CouponEffectiveAt = models.MyTime{Time: stamp}
	user.CouponAmount = data.CouponAmount
	user.CouponProduct = data.CouponProduct
	user.CouponAmountThan = data.CouponAmountThan
	user.Kf = data.Kf
	user.Keywords = data.Keywords
	user.Description = data.Description
	user.Notice = data.Notice
	user.Praise = data.Praise
	user.Rate = data.Rate
	user.MasterLink = data.MasterLink
	user.UserRegCoupon = data.UserRegCoupon
	user.HaulDiscount = data.HaulDiscount
	user.InviteRate = data.InviteRate
	user.WithdrawAmount = data.WithdrawAmount

	if err := ay.Db.Save(&user).Error; err != nil {
		ay.Json{}.Msg(c, 400, "请联系管理员", gin.H{})
		return
	} else {
		ay.Json{}.Msg(c, 200, "修改成功", gin.H{})
		return
	}

}
