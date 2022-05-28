/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type PayModel struct {
}

type Pay struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	Type       int
	Pay        int
	Appid      string
	MchId      string `gorm:"column:mch_id" json:"mch_id"`
	VKey       string `gorm:"column:v_key" json:"v_key"`
	Secret     string `gorm:"column:secret" json:"secret"`
	PayAppid   string `gorm:"column:pay_appid" json:"pay_appid"`
	PayKey     string `gorm:"column:pay_key" json:"pay_key"`
	PayDealId  string `gorm:"column:pay_dealId" json:"pay_dealId"`
	PublicKey  string `gorm:"column:public_key" json:"public_key"`
	PrivateKey string `gorm:"column:private_key" json:"private_key"`
}

func (Pay) TableName() string {
	return "sm_pay"
}
