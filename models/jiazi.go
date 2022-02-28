/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type JiaZiModel struct {
}

type JiaZi struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	JiaZi string `gorm:"column:jiazi" json:"jiazi"`
	NaYin string `gorm:"column:nayin" json:"nayin"`
}

func (JiaZi) TableName() string {
	return "sm_jiazi"
}
