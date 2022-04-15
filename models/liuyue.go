/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type LiuYueModel struct {
}

type LiuYue struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	TianGan string `gorm:"column:tian_gan" json:"tian_gan"`
	YueGan  string `gorm:"column:yue_gan" json:"yue_gan"`
	TenGods string `gorm:"column:ten_gods;FOREIGNKEY:name" json:"ten_gods"`
}

func (LiuYue) TableName() string {
	return "sm_liu_yue"
}
