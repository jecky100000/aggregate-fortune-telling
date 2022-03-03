/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type ConfigModel struct {
}

type Config struct {
	Id                int64   `gorm:"primaryKey" json:"id"`
	HaulAmount        float64 `gorm:"column:haul_amount" json:"haul_amount"`
	Coupon            int     `gorm:"column:coupon" json:"coupon"`
	CouponName        string  `gorm:"column:coupon_name" json:"coupon_name"`
	CouponSubName     string  `gorm:"column:coupon_sub_name" json:"coupon_sub_name"`
	CouponProhibit    string  `gorm:"column:coupon_prohibit" json:"coupon_prohibit"`
	CouponEffective   string  `gorm:"column:coupon_effective" json:"coupon_effective"`
	CouponEffectiveAt MyTime  `gorm:"column:coupon_effective_at" json:"coupon_effective_at"`
	CouponAmount      float64 `gorm:"column:coupon_amount" json:"coupon_amount"`
	CouponProduct     string  `gorm:"column:coupon_product" json:"coupon_product"`
	CouponAmountThan  float64 `gorm:"column:coupon_amount_than" json:"coupon_amount_than"`
	Kf                string  `gorm:"column:kf" json:"kf"`
	Keywords          string  `gorm:"column:keywords" json:"keywords"`
	Description       string  `gorm:"column:description" json:"description"`
	Notice            string  `gorm:"column:notice" json:"notice"`
	Praise            string  `gorm:"column:praise" json:"praise"`
	Rate              float64 `gorm:"column:rate" json:"rate"`
	MasterLink        string  `gorm:"column:master_link" json:"master_link"`
	UserRegCoupon     int     `gorm:"column:user_reg_coupon" json:"user_reg_coupon"`
	HaulDiscount      string  `gorm:"column:haul_discount" json:"haul_discount"`
}

func (Config) TableName() string {
	return "sm_config"
}
