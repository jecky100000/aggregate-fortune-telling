/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

import (
	"gin/ay"
)

type HexapodModel struct {
}

type HexapodList struct {
	Id       int `gorm:"primaryKey"`
	Title    string
	Content  string
	Handount string `gorm:"column:handout"`
}

func (con HexapodModel) GetContonent(title string) HexapodList {
	var ss HexapodList
	ay.Db.Table("sm_hexapod").Where("title", title).Find(&ss)
	return ss
}
