/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type CouponModel struct {
}

type Coupon struct {
	BaseModel
	Uid         int64   `gorm:"column:uid" json:"-"`
	Name        string  `gorm:"column:name" json:"name"`
	SubName     string  `gorm:"column:sub_name" json:"sub_name"`
	Prohibit    string  `gorm:"column:prohibit" json:"prohibit"`
	UsedAt      string  `gorm:"column:used_at" json:"-"`
	Effective   string  `gorm:"column:effective" json:"effective"`
	EffectiveAt MyTime  `gorm:"column:effective_at" json:"effective_at"`
	Amount      float64 `gorm:"column:amount" json:"amount"`
	Product     string  `gorm:"column:product" json:"-"`
	Status      int     `gorm:"column:status" json:"status"`
	AmountThan  float64 `gorm:"column:amount_than" json:"amount_than"`
}

func (Coupon) TableName() string {
	return "sm_user_coupon"
}
