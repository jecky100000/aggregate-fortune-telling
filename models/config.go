/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "gin/ay"

type ConfigModel struct {
}

type Config struct {
	Id                int64   `gorm:"primaryKey" json:"id" form:"id"`
	HaulAmount        float64 `gorm:"column:haul_amount" json:"haul_amount" form:"haul_amount"`
	Coupon            int     `gorm:"column:coupon" json:"coupon" form:"coupon"`
	CouponName        string  `gorm:"column:coupon_name" json:"coupon_name" form:"coupon_name"`
	CouponSubName     string  `gorm:"column:coupon_sub_name" json:"coupon_sub_name" form:"coupon_sub_name"`
	CouponProhibit    string  `gorm:"column:coupon_prohibit" json:"coupon_prohibit" form:"coupon_prohibit"`
	CouponEffective   string  `gorm:"column:coupon_effective" json:"coupon_effective" form:"coupon_effective"`
	CouponEffectiveAt MyTime  `gorm:"column:coupon_effective_at" json:"coupon_effective_at" form:"coupon_effective_at"`
	CouponAmount      float64 `gorm:"column:coupon_amount" json:"coupon_amount" form:"coupon_amount"`
	CouponProduct     string  `gorm:"column:coupon_product" json:"coupon_product" form:"coupon_product"`
	CouponAmountThan  float64 `gorm:"column:coupon_amount_than" json:"coupon_amount_than" form:"coupon_amount_than"`
	Kf                string  `gorm:"column:kf" json:"kf" form:"kf"`
	Keywords          string  `gorm:"column:keywords" json:"keywords" form:"keywords"`
	Description       string  `gorm:"column:description" json:"description" form:"description"`
	Notice            string  `gorm:"column:notice" json:"notice" form:"notice"`
	Praise            string  `gorm:"column:praise" json:"praise" form:"praise"`
	Rate              float64 `gorm:"column:rate" json:"rate" form:"rate"`
	MasterLink        string  `gorm:"column:master_link" json:"master_link" form:"master_link"`
	UserRegCoupon     int     `gorm:"column:user_reg_coupon" json:"user_reg_coupon" form:"user_reg_coupon"`
	HaulDiscount      string  `gorm:"column:haul_discount" json:"haul_discount" form:"haul_discount"`
	InviteRate        float64 `gorm:"column:invite_rate" json:"invite_rate" form:"invite_rate"`
	WithdrawAmount    float64 `gorm:"column:withdraw_amount" json:"withdraw_amount" form:"withdraw_amount"`
}

func (Config) TableName() string {
	return "sm_config"
}

func (con ConfigModel) GetId(id int) (res Config) {
	ay.Db.First(&res, id)
	return
}
